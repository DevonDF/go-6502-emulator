package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type LSRHandler struct {
}

var lsrHandler = &LSRHandler{}

func LSR() *LSRHandler {
	return lsrHandler
}

// shiftRightOneBit shifts the provided value right one bit and returns the new value and whether a carry occured.
func shiftRightOneBit(val uint8) (uint8, bool) {
	lVal := uint16(val)
	result := lVal >> 1
	return uint8(result), (result & 0x100) != 0x00
}

func (handler *LSRHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// 0 -> [] -> C
	var result uint8
	var carry bool

	if instruction.Instruction.AddressingMode == AddrAccumulator {
		byteToShift := cpu.RegisterAC
		result, carry = shiftRightOneBit(uint8(byteToShift))
		cpu.RegisterAC = int8(result)
	} else {
		addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
		if err != nil {
			return err
		}

		byteToShift, err := memory.ReadByte(addr)
		if err != nil {
			return err
		}
		result, carry = shiftRightOneBit(byteToShift)
		memory.Write(addr, []byte{result})
	}

	cpu.SetNegativeFlag(false)
	cpu.SetZeroFlag(result == 0)
	cpu.SetCarryFlag(carry)
	return nil
}
