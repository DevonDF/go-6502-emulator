package emulator

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
	"golang.org/x/term"
)

type EmulatorConfiguration struct {
	Debug bool
}

type Emulator struct {
	cpu     *cpu.CPU
	memory  *memory.Memory
	config  EmulatorConfiguration
	running bool
	logger  *slog.Logger
}

// NewEmulator creates a new Emulator with the provided configuration
func NewEmulator(config EmulatorConfiguration) *Emulator {
	loggerLevel := slog.LevelError
	if config.Debug {
		loggerLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: loggerLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch v := a.Value.Any().(type) {
			case int:
				a.Value = slog.StringValue(fmt.Sprintf("0x%X", v))
			case int8:
				a.Value = slog.StringValue(fmt.Sprintf("0x%02X", uint8(v)))
			case int16:
				a.Value = slog.StringValue(fmt.Sprintf("0x%04X", uint16(v)))
			case int32:
				a.Value = slog.StringValue(fmt.Sprintf("0x%08X", uint32(v)))
			case int64:
				a.Value = slog.StringValue(fmt.Sprintf("0x%X", v))
			case uint, uint8, uint16, uint32, uint64:
				a.Value = slog.StringValue(fmt.Sprintf("0x%X", a.Value.Uint64()))
			case []byte:
				a.Value = slog.StringValue(fmt.Sprintf("0x%s", hex.EncodeToString(v)))
			}
			return a
		},
	}))

	return &Emulator{
		cpu:     cpu.NewCPU(logger),
		memory:  memory.NewMemory(logger),
		config:  config,
		running: false,
		logger:  logger,
	}
}

// StartEmulator runs any startup required when beginning the emulator and starts the interactive emulator.
func (emulator *Emulator) StartEmulator() {
	emulator.cpu.RegisterPC = memory.ROMStartAddress // Set the PC to the start of ROM
	emulator.cpu.RegisterSP = 0x00                   // Set stack pointer to 0x00
	emulator.cpu.RegisterX = 0x00                    // Set register X to 0x00
	emulator.cpu.RegisterY = 0x00                    // Set register Y to 0x00

	fd := int(os.Stdin.Fd())

	// Put terminal into raw mode
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	// Ensure terminal is restored on Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		term.Restore(fd, oldState)
		os.Exit(0)
	}()

	// Create an inputChannel and start a goroutine to populate it with user input.
	inputChannel := make(chan byte)
	go func() {
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err == nil {
				inputChannel <- buf[0]
			}
		}
	}()

	// Start the main thread of the emulator.
	running := false
	emulator.printInteractiveTerminal()
	for {
		select {
		case key := <-inputChannel:
			// an input was detected by the user
			switch key {
			case 'q':
				return
			case 's':
				err := emulator.step()
				if err != nil {
					return
				}
				emulator.printInteractiveTerminal()
			case 'r':
				running = true
			}
		default:
			if running {
				err := emulator.step()
				if err != nil {
					return
				}
				emulator.printInteractiveTerminal()
				time.Sleep(time.Millisecond * 100)
			}
		}
	}
}

// stopEmulator stops the emulator and resets the state if required.
func (emulator *Emulator) StopEmulator() {
	emulator.running = false

	if emulator.config.Debug {
		emulator.logger.Debug("dumping memory to memory.dump")
		os.WriteFile("memory.dump", []byte(emulator.memory.Dump()), 0644)
	}
}

// LoadROM loads a ROM into memory at 0x8000.
func (emulator *Emulator) LoadROM(romPath string) error {
	emulator.logger.Debug("loading ROM", "rom", romPath)
	rom, err := os.ReadFile(romPath)
	if err != nil {
		return fmt.Errorf("failed to load ROM %s", romPath)
	}
	emulator.memory.Write(memory.ROMStartAddress, rom)
	emulator.logger.Debug("wrote ROM into memory", "size", len(rom))
	return nil
}

// generateDescriptiveString generates a descriptive string of the emulator's state.
func (emulator *Emulator) generateDescriptiveString() string {
	currentInstruction, err := emulator.fetchAndDecodeCurrentInstruction()
	if err != nil {
		return ""
	}

	str := "Registers\n"
	str += fmt.Sprintf("\tPC = 0x%04X <- %s\n", emulator.cpu.RegisterPC, currentInstruction.GetFullAssemblyString())
	str += fmt.Sprintf("\tAC = 0x%02X\n", emulator.cpu.RegisterAC)
	str += fmt.Sprintf("\tX = 0x%02X\n", emulator.cpu.RegisterX)
	str += fmt.Sprintf("\tY = 0x%02X\n", emulator.cpu.RegisterY)
	str += fmt.Sprintf("\tSP = 0x%02X\n", emulator.cpu.RegisterSP)
	str += fmt.Sprintf("\tSR = 0x%02X    C=0x%02X    Z=0x%02X    I=0x%02X    D=0x%02X    B=0x%02X    V=0x%02X    N=0x%02X\n\n",
		emulator.cpu.RegisterSR,
		emulator.cpu.GetCarryFlag(),
		emulator.cpu.GetZeroFlag(),
		emulator.cpu.GetInterruptFlag(),
		emulator.cpu.GetDecimalFlag(),
		emulator.cpu.GetBreakFlag(),
		emulator.cpu.GetOverflowFlag(),
		emulator.cpu.GetNegativeFlag())

	str += "Stack\n"
	str += emulator.cpu.Stack.ToHexDump(emulator.memory) + "\n"

	str += "Zeropage\n"
	zeropageBytes := make([]byte, 0x100)
	emulator.memory.Read(0x00, &zeropageBytes)
	str += hex.Dump(zeropageBytes) + "\n"

	return str
}

func (emulator *Emulator) printInteractiveTerminal() {
	fmt.Print("\033[2J\033[H")
	fmt.Print("6502 Interactive Terminal Emulator\n\n")
	fmt.Print(emulator.generateDescriptiveString())
	fmt.Print("\ns=step    r=run    p=pause    q=quit\n")
}

// fetchAndDecodeCurrentInstruction fetches and decodes the current instruction pointed to by PC.
func (emulator *Emulator) fetchAndDecodeCurrentInstruction() (*instructions.DecodedInstruction, error) {
	// Fetch
	reader, err := emulator.memory.ReaderAt(emulator.cpu.RegisterPC)
	if err != nil {
		emulator.logger.Error("failed to fetch instruction from memory", "error", err)
		return nil, err
	}

	// Decode
	inst, err := instructions.DecodeNextInstruction(reader)
	if err != nil {
		emulator.logger.Error("failed to decode next instruction", "error", err)
		return nil, err
	}
	return inst, nil
}

// Step performs one fetch-decode-execute loop within the CPU.
func (emulator *Emulator) step() error {
	// fetch and decode the current instruction on PC
	inst, err := emulator.fetchAndDecodeCurrentInstruction()
	if err != nil {
		return err
	}

	// Execute
	emulator.logger.Debug("executing instruction", "opcode", inst.Instruction.Opcode, "instruction", inst.Instruction.AssemblyString, "operands", inst.Operands)
	err = inst.Instruction.Handler.Execute(emulator.cpu, emulator.memory, inst)
	if err != nil {
		emulator.logger.Error("execution of instruction caused error", "opcode", inst.Instruction.Opcode, "err", err)
		return err
	}

	// Increment PC
	emulator.cpu.RegisterPC += uint16(inst.Instruction.Size)

	return nil
}
