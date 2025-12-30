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
	cpu.RegisterSR = cpu.RegisterSR | 0x01
}

// SetZeroFlag sets the zero bit flag.
func (cpu *CPU) SetZeroFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x02
}

// SetInterruptFlag sets the interrupt bit flag.
func (cpu *CPU) SetInterruptFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x04
}

// SetDecimalFlag sets the decimal bit flag.
func (cpu *CPU) SetDecimalFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x08
}

// SetBreakFlag sets the break bit flag.
func (cpu *CPU) SetBreakFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x10
}

// SetOverflowFlag sets the overflow bit flag.
func (cpu *CPU) SetOverflowFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x40
}

// SetNegativeFlag sets the negative bit flag.
func (cpu *CPU) SetNegativeFlag(bit bool) {
	cpu.RegisterSR = cpu.RegisterSR | 0x80
}

// SetStatusFlags sets the flags for the SR Status Register within the CPU.
func (cpu *CPU) SetStatusFlags(negative bool, overflow bool, break_ bool, decimal bool, interrupt bool, zero bool, carry bool) {
	newSR := uint8(0)

	if negative {
		newSR = newSR | 0x80
	}
	if overflow {
		newSR = newSR | 0x40
	}
	if break_ {
		newSR = newSR | 0x10
	}
	if decimal {
		newSR = newSR | 0x08
	}
	if interrupt {
		newSR = newSR | 0x04
	}
	if zero {
		newSR = newSR | 0x02
	}
	if carry {
		newSR = newSR | 0x01
	}

	cpu.RegisterSR = newSR
}

// LogRegisters logs the state of the registers to the debug logger.
func (cpu *CPU) LogRegisters() {
	cpu.logger.Debug("CPU registers", "pc", cpu.RegisterPC, "ac", cpu.RegisterAC, "x", cpu.RegisterX, "y", cpu.RegisterY, "sp", cpu.RegisterSP, "sr", cpu.RegisterSR)
}
