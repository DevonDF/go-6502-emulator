package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type DECHandler struct {
}

var decHandler = &DECHandler{}

func DEC() *DECHandler {
	return decHandler
}

func (handler *DECHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// M <- M - 1
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	readByte, err := memory.ReadByte(addr)
	if err != nil {
		return err
	}

	result := int8(readByte) - 1
	memory.Write(addr, []byte{byte(result)})

	cpu.SetNegativeFlag(result < 0)
	cpu.SetZeroFlag(result == 0)
	return nil
}
