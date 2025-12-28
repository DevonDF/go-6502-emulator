package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type ADCHandler struct {
}

var adcHandler = &ADCHandler{}

func ADC() *ADCHandler {
	return adcHandler
}

func (handler *ADCHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {

	var valueToAdd int8

	switch instruction.Instruction.Opcode {

	case 0x69: // immediate	ADC #oper
		valueToAdd = int8(instruction.Operands[0])

	case 0x65: // zeropage	ADC oper
		addr := uint8(instruction.Operands[0])
		readByte, err := memory.ReadByte(uint16(addr))
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x75: // zeropage,X	ADC oper,X
		addr := uint8(instruction.Operands[0])
		readByte, err := memory.ReadByte(uint16(addr + uint8(cpu.RegisterX)))
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x6D: // absolute	ADC oper
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr)
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x7D: // absolute,X	ADC oper,X
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr + uint16(cpu.RegisterX))
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x79: // absolute,Y	ADC oper,Y
		addr := uint16(instruction.Operands[1])<<8 | uint16(instruction.Operands[0])
		readByte, err := memory.ReadByte(addr + uint16(cpu.RegisterY))
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x61: // (indirect,X)	ADC (oper,X)
		addr := uint16(instruction.Operands[0]) + uint16(cpu.RegisterX)
		addr2, err := memory.Read16(addr)
		if err != nil {
			return err
		}
		readByte, err := memory.ReadByte(addr2)
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)

	case 0x71: // (indirect),Y	ADC (oper),Y
		addr := uint16(instruction.Operands[0])
		addr2, err := memory.Read16(addr)
		if err != nil {
			return err
		}
		readByte, err := memory.ReadByte(addr2 + uint16(cpu.RegisterY))
		if err != nil {
			return err
		}
		valueToAdd = int8(readByte)
	}

	cpu.Accumulator.Add(int8(valueToAdd))
	return nil
}
