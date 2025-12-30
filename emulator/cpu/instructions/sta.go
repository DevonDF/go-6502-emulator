package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type STAHandler struct {
}

var staHandler = &STAHandler{}

func STA() *STAHandler {
	return staHandler
}

func (handler *STAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	memory.Write(addr, []byte{byte(cpu.RegisterAC)})
	return nil
}
