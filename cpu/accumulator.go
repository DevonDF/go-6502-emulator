package cpu

import "log/slog"

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
	rAC := int16(acu.cpu.rAC)
	num := int16(number)
	carryIn := int16(acu.cpu.GetCarryFlag())

	sum := rAC + num + carryIn
	result := int8(sum)

	// if the sum is over 127 or below -128 a carry occured (too large to fit in a signed byte)
	carryOut := sum > 0x7F || sum < -0x80

	// if the 8th bit is set incorrectly (i.e. a signed overflow occured) then the overflow flag should be set
	overflowOut := int16(acu.cpu.rAC^result)&int16(number^result)&0x80 != 0x0

	acu.cpu.rAC = result
	acu.cpu.SetStatusFlags(result < 0, overflowOut, false, false, false, result == 0, carryOut)

	acu.logger.Debug("acu.Add", "rAC", rAC, "operand", num, "carryIn", carryIn, "result", result, "carryOut", carryOut, "overflowOut", overflowOut)
}
