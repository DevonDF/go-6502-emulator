package emulator

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
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

// startEmulator runs any startup required when beginning the emulator.
func (emulator *Emulator) StartEmulator() {
	emulator.logger.Debug("starting emulator")
	emulator.running = true
	emulator.cpu.RegisterPC = memory.ROMStartAddress // Set the PC to the start of ROM
	emulator.cpu.RegisterSP = 0x00                   // Set stack pointer to 0x00
	emulator.cpu.RegisterX = 0x00                    // Set register X to 0x00
	emulator.cpu.RegisterY = 0x00                    // Set register Y to 0x00
	emulator.execute()
}

// stopEmulator stops the emulator and resets the state if required.
func (emulator *Emulator) StopEmulator() {
	emulator.logger.Debug("stopping emulator")
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

// execute begins the fetch-decode-execute loop at the PC.
func (emulator *Emulator) execute() {
	defer emulator.StopEmulator()

	// Start fetch-decode-execute loop
	for emulator.running {
		// Fetch
		reader, err := emulator.memory.ReaderAt(emulator.cpu.RegisterPC)
		if err != nil {
			emulator.logger.Error("failed to fetch instruction from memory", "error", err)
		}

		// Decode
		inst, err := instructions.DecodeNextInstruction(reader)
		if err != nil {
			emulator.logger.Error("failed to decode next instruction", "error", err)
			return
		}

		// Execute
		emulator.logger.Debug("executing instruction", "opcode", inst.Instruction.Opcode, "instruction", inst.Instruction.AssemblyString, "operands", inst.Operands)
		err = inst.Instruction.Handler.Execute(emulator.cpu, emulator.memory, &inst)
		if err != nil {
			emulator.logger.Error("execution of instruction caused error", "opcode", inst.Instruction.Opcode, "err", err)
			return
		}

		// Increment PC
		emulator.cpu.RegisterPC += uint16(inst.Instruction.Size)
		emulator.cpu.LogRegisters()
	}

}
