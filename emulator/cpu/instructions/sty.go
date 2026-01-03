package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type STYHandler struct {
}

var styHandler = &STYHandler{}

func STY() *STYHandler {
	return styHandler
}

func (handler *STYHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	memory.Write(addr, []byte{byte(cpu.RegisterY)})
	return nil
}
