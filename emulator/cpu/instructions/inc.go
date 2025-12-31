package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type INCHandler struct {
}

var incHandler = &INCHandler{}

func INC() *INCHandler {
	return incHandler
}

func (handler *INCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// M + 1 -> M
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	readByte, err := memory.ReadByte(addr)
	if err != nil {
		return err
	}

	result := int8(readByte) + 1
	memory.Write(addr, []byte{byte(result)})

	cpu.SetNegativeFlag(result < 0)
	cpu.SetZeroFlag(result == 0)
	return nil
}
