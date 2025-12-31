package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type ROLHandler struct {
}

var rolHandler = &ROLHandler{}

func ROL() *ROLHandler {
	return rolHandler
}

// rotateLeftOneBit rotates the provided value left one bit and returns the new value and whether a carry occured.
func rotateLeftOneBit(val uint8) (uint8, bool) {
	lVal := uint16(val)
	result := lVal << 1
	result = result | ((result >> 8) & 0x1)
	return uint8(result), (result & 0x100) != 0x00
}

func (handler *ROLHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// C <- [] <- C
	var result uint8
	var carry bool

	if instruction.Instruction.AddressingMode == AddrAccumulator {
		byteToShift := cpu.RegisterAC
		result, carry = rotateLeftOneBit(uint8(byteToShift))
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
		result, carry = rotateLeftOneBit(byteToShift)
		memory.Write(addr, []byte{result})
	}

	cpu.SetNegativeFlag(int8(result) < 0)
	cpu.SetZeroFlag(result == 0)
	cpu.SetCarryFlag(carry)
	return nil
}
