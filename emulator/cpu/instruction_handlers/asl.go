package instruction_handlers

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/cpu/addressing"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ASLHandler struct {
}

var aslHandler = &ASLHandler{}

func ASL() *ASLHandler {
	return aslHandler
}

// shiftLeftOneBit shifts the provided value left one bit and returns the new value and whether a carry occured.
func shiftLeftOneBit(val uint8) (uint8, bool) {
	lVal := uint16(val)
	result := lVal << 1
	return uint8(result), (result & 0x100) != 0x00
}

func (handler *ASLHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	var addr uint16
	var err error

	switch instruction.Instruction.Opcode {
	case 0x0A: // accumulator	ASL A
		cpu.Accumulator.ShiftLeftOneBit()
		return nil // all done within accumulator, should stop here

	case 0x06: // zeropage	ASL oper
		addr, err = addressing.GetZeropageAddress(instruction.Operands, cpu, memory)

	case 0x16: // zeropage,X	ASL oper,X
		addr, err = addressing.GetZeropageXAddress(instruction.Operands, cpu, memory)

	case 0x0E: // absolute	ASL oper
		addr, err = addressing.GetAbsoluteAddress(instruction.Operands, cpu, memory)

	case 0x1E: // absolute,X	ASL oper,X
		addr, err = addressing.GetAbsoluteXAddress(instruction.Operands, cpu, memory)
	}

	readByte, err := memory.ReadByte(addr)
	if err != nil {
		return err
	}
	result, carry := shiftLeftOneBit(readByte)
	memory.Write(addr, []byte{result})

	cpu.SetStatusFlags(int8(result) < 0, false, false, false, false, result == 0, carry)
	return nil
}
