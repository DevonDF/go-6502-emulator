package cpu

import "log/slog"

// Accumulator handles mathematical & bitwise operations within the ACU unit and AC register.
type Accumulator struct {
	cpu    *CPU
	logger *slog.Logger
}

// NewAccumulator creates a new ACU for a given CPU.
func NewAccumulator(cpu *CPU, logger *slog.Logger) *Accumulator {
	return &Accumulator{
		cpu:    cpu,
		logger: logger,
	}
}

// Add adds a given number to the accumulator and sets any required registers and bits accordingly.
func (acu *Accumulator) Add(number int8) {
	// cast these to a larger space, and then analyse the result to
	// get the carry & overflow flags
	rAC := int16(acu.cpu.RegisterAC)
	num := int16(number)
	carryIn := int16(acu.cpu.GetCarryFlag())

	sum := rAC + num + carryIn
	result := int8(sum)

	// if the sum is over 127 or below -128 a carry occured (too large to fit in a signed byte)
	carryOut := sum > 0x7F || sum < -0x80

	// if the 8th bit is set incorrectly (i.e. a signed overflow occured) then the overflow flag should be set
	overflowOut := int16(acu.cpu.RegisterAC^result)&int16(number^result)&0x80 != 0x0

	acu.cpu.RegisterAC = result
	acu.cpu.SetNegativeFlag(result < 0)
	acu.cpu.SetZeroFlag(result == 0)
	acu.cpu.SetCarryFlag(carryOut)

	acu.logger.Debug("acu.Add", "rAC", rAC, "operand", num, "carryIn", carryIn, "result", result, "carryOut", carryOut, "overflowOut", overflowOut)
}

// And performs a bitwise logic and within the accumulator with the provided operand.
func (acu *Accumulator) And(operand uint8) {
	rAC := uint8(acu.cpu.RegisterAC)
	result := int8(rAC & operand)
	acu.cpu.RegisterAC = result
	acu.cpu.SetNegativeFlag(result < 0)
	acu.cpu.SetZeroFlag(result == 0)
	acu.logger.Debug("acu.And", "rAC", rAC, "operand", operand, "result", result)
}

// ShiftLeftOneBit shifts the value of AC one bit.
func (acu *Accumulator) ShiftLeftOneBit() {
	// cast to larger space so we can observe carry
	rAC := uint16(acu.cpu.RegisterAC)
	result := rAC << 1
	carryOut := (result & 0x100) != 0x00

	acu.cpu.RegisterAC = int8(result)
	acu.cpu.SetNegativeFlag(int8(result) < 0)
	acu.cpu.SetZeroFlag(result == 0)
	acu.cpu.SetCarryFlag(carryOut)
	acu.logger.Debug("acu.ShiftLeftOneBit", "rAC", rAC, "result", result, "carryOut", carryOut)
}
