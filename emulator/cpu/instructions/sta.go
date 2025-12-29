package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type STAHandler struct {
}

var staHandler = &STAHandler{}

func STA() *STAHandler {
	return staHandler
}

func (handler *STAHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	var addr uint16
	var err error

	switch instruction.Instruction.Opcode {
	case 0x85: // zeropage	STA oper
		addr, err = addressing.GetZeropageAddress(instruction.Operands, cpu, memory)

	case 0x95: // zeropage,X	STA oper,X
		addr, err = addressing.GetZeropageXAddress(instruction.Operands, cpu, memory)

	case 0x8D: // absolute	STA oper
		addr, err = addressing.GetAbsoluteAddress(instruction.Operands, cpu, memory)

	case 0x9D: // absolute,X	STA oper,X
		addr, err = addressing.GetAbsoluteXAddress(instruction.Operands, cpu, memory)

	case 0x99: // absolute,Y	STA oper,Y
		addr, err = addressing.GetAbsoluteYAddress(instruction.Operands, cpu, memory)

	case 0x81: // (indirect,X)	STA (oper,X)
		addr, err = addressing.GetIndirectXAddress(instruction.Operands, cpu, memory)

	case 0x91: // (indirect),Y	STA (oper),Y
		addr, err = addressing.GetIndirectYAddress(instruction.Operands, cpu, memory)
	}

	if err != nil {
		return err
	}
	memory.Write(addr, []byte{byte(cpu.RegisterAC)})
	return nil
}
