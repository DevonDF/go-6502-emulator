package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type JMPHandler struct {
}

var jmpHandler = &JMPHandler{}

func JMP() *JMPHandler {
	return jmpHandler
}

func (handler *JMPHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// M -> PC

	var nextPC uint16
	var err error

	switch instruction.Instruction.AddressingMode {
	case AddrAbsolute:
		nextPC, err = addressing.GetAbsoluteAddress(instruction.Operands, cpu, memory)
	case AddrIndirect:
		addr, err := addressing.GetAbsoluteAddress(instruction.Operands, cpu, memory)
		if err != nil {
			return err
		}
		nextPC, err = memory.ReadDouble(addr)
	}

	if err != nil {
		return err
	}

	cpu.RegisterPC = nextPC
	return nil
}
