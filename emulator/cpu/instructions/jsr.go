package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type JSRHandler struct {
}

var jsrHandler = &JSRHandler{}

func JSR() *JSRHandler {
	return jsrHandler
}

func (handler *JSRHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// push (PC+2)
	// M -> PC
	cpu.Stack.PushDouble(int16(cpu.RegisterPC)+2, memory)

	nextPC, err := addressing.GetAbsoluteAddress(instruction.Operands, cpu, memory)
	if err != nil {
		return err
	}

	// the fetch-decode-execute cycle will increment by the instruction size, we should counteract this
	cpu.RegisterPC = (nextPC - uint16(instruction.Instruction.Size))
	return nil
}
