package emulator

import (
	"bufio"
	"emulator/cpu"
	"emulator/cpu/instructions"
	"emulator/memory"
	"log/slog"
	"os"
)

type EmulatorConfiguration struct {
	Verbose bool
}

type Emulator struct {
	cpu    *cpu.CPU
	memory *memory.Memory
	config EmulatorConfiguration
	logger *slog.Logger
}

// NewEmulator creates a new Emulator with the provided configuration
func NewEmulator(config EmulatorConfiguration) *Emulator {
	loggerLevel := slog.LevelError
	if config.Verbose {
		loggerLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: loggerLevel,
	}))

	return &Emulator{
		cpu:    cpu.NewCPU(logger),
		memory: memory.NewMemory(logger),
		config: config,
		logger: logger,
	}
}

// startEmulator runs any startup required when beginning the emulator.
func (emulator *Emulator) startEmulator() {
	emulator.logger.Debug("starting emulator")
}

// stopEmulator stops the emulator and resets the state if required.
func (emulator *Emulator) stopEmulator() {
	emulator.logger.Debug("stopping emulator")
}

// LoadAndExecute loads a given ROM at the provided path and executes it within the emulator.
func (emulator *Emulator) LoadAndExecute(romPath string) {
	defer emulator.stopEmulator()
	emulator.startEmulator()

	// Open ROM in a buffered reader
	emulator.logger.Debug("loading ROM", "rom", romPath)
	rom, err := os.Open(romPath)
	if err != nil {
		emulator.logger.Error("failed to load ROM", "rom", romPath)
		return
	}
	defer rom.Close()

	reader := bufio.NewReader(rom)

	// Start fetch-decode-execute loop
	for true {

		// Fetch & decode next instruction
		inst, err := instructions.DecodeNextInstruction(reader)
		if err != nil {
			emulator.logger.Error("failed to fetch and decode next instruction", "error", err)
			return
		}
		emulator.logger.Debug("executing instruction", "opcode", inst.Instruction.Opcode, "operands", inst.Operands)
		inst.Instruction.Handler.Execute(emulator.cpu, emulator.memory, &inst)

		emulator.cpu.LogState()
	}

}
