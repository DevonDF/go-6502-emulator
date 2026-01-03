package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type STXHandler struct {
}

var stxHandler = &STXHandler{}

func STX() *STXHandler {
	return stxHandler
}

func (handler *STXHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	memory.Write(addr, []byte{byte(cpu.RegisterX)})
	return nil
}
