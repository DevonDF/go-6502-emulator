package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type ANDHandler struct {
}

var andHandler = &ANDHandler{}

func AND() *ANDHandler {
	return andHandler
}

func (handler *ANDHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {

	var value int8

	switch instruction.Instruction.Opcode {

	case 0x29: // immediate	AND #oper
		value = int8(instruction.Operands[0])

	case 0x25: // zeropage	AND oper
		addr := uint8(instruction.Operands[0])
		readByte, err := memory.ReadByte(uint16(addr))
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x35: // zeropage,X	AND oper,X
		addr := uint8(instruction.Operands[0])
		readByte, err := memory.ReadByte(uint16(addr + uint8(cpu.RegisterX)))
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x2D: // absolute	AND oper
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr)
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x3D: // absolute,X	AND oper,X
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr + uint16(cpu.RegisterX))
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x39: // absolute,Y	AND oper,Y
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr + uint16(cpu.RegisterY))
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x21: // (indirect,X)	AND (oper,X)
		addr := uint16(instruction.Operands[0]) + uint16(cpu.RegisterX)
		addr2, err := memory.Read16(addr)
		if err != nil {
			return err
		}
		readByte, err := memory.ReadByte(addr2)
		if err != nil {
			return err
		}
		value = int8(readByte)

	case 0x31: // (indirect),Y	AND (oper),Y
		addr := uint16(instruction.Operands[0])
		addr2, err := memory.Read16(addr)
		if err != nil {
			return err
		}
		readByte, err := memory.ReadByte(addr2 + uint16(cpu.RegisterY))
		if err != nil {
			return err
		}
		value = int8(readByte)
	}

	cpu.Accumulator.And(uint8(value))
	return nil
}
