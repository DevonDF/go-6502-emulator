package cpu

import (
	"log/slog"
)

// CPU holds the CPU registers and an ACU for mathematical & bitwise operations.
type CPU struct {
	Accumulator *Accumulator
	RegisterPC  uint16
	RegisterAC  int8
	RegisterX   int8
	RegisterY   int8
	RegisterSR  uint8
	RegisterSP  int8

	logger *slog.Logger
}

// NewCPU creates a new CPU struct and returns it.
func NewCPU(logger *slog.Logger) *CPU {
	cpu := &CPU{
		RegisterPC: 0,
		RegisterAC: 0,
		RegisterX:  0,
		RegisterY:  0,
		RegisterSR: 0,
		RegisterSP: 0,

		logger: logger,
	}
	cpu.Accumulator = NewAccumulator(cpu, logger)
	return cpu
}

// GetCarryFlag returns the carry bit flag.
func (cpu *CPU) GetCarryFlag() uint8 {
	return cpu.RegisterSR & 0x01
}

// GetZeroFlag returns the zero bit flag.
func (cpu *CPU) GetZeroFlag() uint8 {
	return (cpu.RegisterSR & 0x02) >> 1
}

// GetInterruptFlag returns the interrupt bit flag.
func (cpu *CPU) GetInterruptFlag() uint8 {
	return (cpu.RegisterSR & 0x04) >> 2
}

// GetDecimalFlag returns the decimal bit flag.
func (cpu *CPU) GetDecimalFlag() uint8 {
	return (cpu.RegisterSR & 0x08) >> 3
}

// GetBreakFlag returns the break bit flag.
func (cpu *CPU) GetBreakFlag() uint8 {
	return (cpu.RegisterSR & 0x10) >> 4
}

// GetOverflowFlag returns the overflow bit flag.
func (cpu *CPU) GetOverflowFlag() uint8 {
	return (cpu.RegisterSR & 0x40) >> 6
}

// GetNegativeFlag returns the negative bit flag.
func (cpu *CPU) GetNegativeFlag() uint8 {
	return (cpu.RegisterSR & 0x80) >> 7
}

// SetCarryFlag sets the carry bit flag.
func (cpu *CPU) SetCarryFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x01
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x01)
	}

}

// SetZeroFlag sets the zero bit flag.
func (cpu *CPU) SetZeroFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x02
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x02)
	}
}

// SetInterruptFlag sets the interrupt bit flag.
func (cpu *CPU) SetInterruptFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x04
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x04)
	}
}

// SetDecimalFlag sets the decimal bit flag.
func (cpu *CPU) SetDecimalFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x08
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x08)
	}
}

// SetBreakFlag sets the break bit flag.
func (cpu *CPU) SetBreakFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x10
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x10)
	}
}

// SetOverflowFlag sets the overflow bit flag.
func (cpu *CPU) SetOverflowFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x40
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x40)
	}
}

// SetNegativeFlag sets the negative bit flag.
func (cpu *CPU) SetNegativeFlag(bit bool) {
	if bit {
		cpu.RegisterSR = cpu.RegisterSR | 0x80
	} else {
		cpu.RegisterSR = cpu.RegisterSR & (0xFF - 0x80)
	}
}

// LogRegisters logs the state of the registers to the debug logger.
func (cpu *CPU) LogRegisters() {
	cpu.logger.Debug("CPU registers", "pc", cpu.RegisterPC, "ac", cpu.RegisterAC, "x", cpu.RegisterX,
		"y", cpu.RegisterY, "sp", cpu.RegisterSP, "sr", cpu.RegisterSR, "N", cpu.GetNegativeFlag(), "V",
		cpu.GetOverflowFlag(), "B", cpu.GetBreakFlag(), "D", cpu.GetDecimalFlag(), "I", cpu.GetInterruptFlag(),
		"Z", cpu.GetZeroFlag(), "C", cpu.GetCarryFlag())
}
