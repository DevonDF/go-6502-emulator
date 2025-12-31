package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type PLPHandler struct {
}

var plpHandler = &PLPHandler{}

func PLP() *PLPHandler {
	return plpHandler
}

func (handler *PLPHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// pull SR
	// The status register will be pulled with the break flag and bit 5 ignored. TODO
	poppedByte, err := cpu.Stack.PopByte(memory)
	if err != nil {
		return err
	}
	cpu.RegisterSR = uint8(poppedByte)
	return nil
}
