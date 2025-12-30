package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
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
	addr, err := instruction.GetOperandReferencedAddress(cpu, memory)
	if err != nil {
		return err
	}

	readByte, err := memory.ReadByte(addr)
	if err != nil {
		return err
	}
	result, carry := shiftLeftOneBit(readByte)
	memory.Write(addr, []byte{result})

	cpu.SetNegativeFlag(int8(result) < 0)
	cpu.SetZeroFlag(result == 0)
	cpu.SetCarryFlag(carry)
	return nil
}
