package instruction_handlers

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type BMIHandler struct {
}

var bmiHandler = &BMIHandler{}

func BMI() *BMIHandler {
	return bmiHandler
}

func (handler *BMIHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// check if N = 1
	if cpu.GetNegativeFlag() == 1 {
		// jump relative
		// -2 as the fetch-decode-execute cycle will increment the PC by 2 after this instruction
		cpu.RegisterPC += (uint16(instruction.Operands[0]) - 2)
	}
	return nil
}
