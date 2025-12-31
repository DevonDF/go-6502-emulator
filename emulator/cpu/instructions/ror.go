package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type RORHandler struct {
}

var rorHandler = &RORHandler{}

func ROR() *RORHandler {
	return rorHandler
}

// rotateRightOneBit rotates the provided value right one bit and returns the new value and whether a carry occured.
func rotateRightOneBit(val uint8) (uint8, bool) {
	lVal := uint16(val)
	endBit := lVal & 0x1
	result := lVal >> 1
	result = result | (endBit << 7)
	return uint8(result), endBit == 0x1
}

func (handler *RORHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// C -> [] -> C
	var result uint8
	var carry bool

	if instruction.Instruction.AddressingMode == AddrAccumulator {
		byteToShift := cpu.RegisterAC
		result, carry = rotateRightOneBit(uint8(byteToShift))
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
		result, carry = rotateRightOneBit(byteToShift)
		memory.Write(addr, []byte{result})
	}

	cpu.SetNegativeFlag(int8(result) < 0)
	cpu.SetZeroFlag(result == 0)
	cpu.SetCarryFlag(carry)
	return nil
}
