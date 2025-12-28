package cpu

import (
	"log/slog"
)

type CPU struct {
	Accumulator *Accumulator

	rPC    int16
	rAC    int8
	rX     int8
	rY     int8
	rSR    uint8
	rSP    int8
	logger *slog.Logger
}

// NewCPU creates a new CPU struct and returns it.
func NewCPU(logger *slog.Logger) *CPU {
	cpu := &CPU{
		rPC: 0,
		rAC: 0,
		rX:  0,
		rY:  0,
		rSR: 0,
		rSP: 0,

		logger: logger,
	}
	cpu.Accumulator = NewAccumulator(cpu, logger)
	return cpu
}

// GetCarryFlag returns the carry bit flag.
func (cpu *CPU) GetCarryFlag() uint8 {
	return cpu.rSR & 0x01
}

// GetZeroFlag returns the zero bit flag.
func (cpu *CPU) GetZeroFlag() uint8 {
	return (cpu.rSR & 0x02) >> 1
}

// GetInterruptFlag returns the interrupt bit flag.
func (cpu *CPU) GetInterruptFlag() uint8 {
	return (cpu.rSR & 0x04) >> 2
}

// GetDecimalFlag returns the decimal bit flag.
func (cpu *CPU) GetDecimalFlag() uint8 {
	return (cpu.rSR & 0x08) >> 3
}

// GetBreakFlag returns the break bit flag.
func (cpu *CPU) GetBreakFlag() uint8 {
	return (cpu.rSR & 0x10) >> 4
}

// GetOverflowFlag returns the overflow bit flag.
func (cpu *CPU) GetOverflowFlag() uint8 {
	return (cpu.rSR & 0x40) >> 6
}

// GetNegativeFlag returns the negative bit flag.
func (cpu *CPU) GetNegativeFlag() uint8 {
	return (cpu.rSR & 0x80) >> 7
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

	cpu.rSR = newSR
}

func (cpu *CPU) LogState() {
	cpu.logger.Debug("CPU State", "rPC", cpu.rPC, "rAC", cpu.rAC, "rX", cpu.rX, "rY", cpu.rY, "rSR", cpu.rSR, "rSP", cpu.rSP)
}
