package assembler

import (
	"fmt"
	"strconv"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type EncodedInstruction struct {
	assembly        string                    // the raw assembly for this instruction.
	operands        string                    // the operands string.
	instruction     *instructions.Instruction // the instruction of this instruction.
	relativeAddress int                       // the relative address of this instruction.
}

// ToBytecode compiles and returns into bytecode the instruction.
func (encodedInstruction *EncodedInstruction) ToBytecode(assembler *Assembler) ([]byte, error) {
	var operandBytes []byte

	switch encodedInstruction.instruction.AddressingMode {
	case instructions.AddrNone:
		operandBytes = []byte{}

	case instructions.AddrAccumulator:
		operandBytes = []byte{}

	case instructions.AddrRelative:
		labelAddr, found := assembler.getAddressForLabel(encodedInstruction.operands)
		if !found {
			return nil, fmt.Errorf("unable to find label %s: %s", encodedInstruction.operands, encodedInstruction.assembly)
		}
		operandBytes = []byte{byte((labelAddr - encodedInstruction.relativeAddress) - int(encodedInstruction.instruction.Size))}

	case instructions.AddrImmediate:
		if encodedInstruction.operands[1] != '&' {
			return nil, fmt.Errorf("invalid immediate addressing used: %s", encodedInstruction.operands)
		}
		v, err := strconv.ParseUint(encodedInstruction.operands[2:], 16, 8)
		if err != nil {
			return nil, err
		}
		operandBytes = []byte{byte(v)}

	case instructions.AddrZeropage:
		v, err := strconv.ParseUint(encodedInstruction.operands[1:], 16, 8)
		if err != nil {
			return nil, err
		}
		operandBytes = []byte{byte(v)}

	case instructions.AddrAbsolute:
		v, err := strconv.ParseUint(encodedInstruction.operands[1:], 16, 8)
		if err != nil {
			return nil, err
		}
		operandBytes = []byte{byte((v >> 8) & 0xFF), byte(v & 0xFF)}

	default:
		return nil, fmt.Errorf("unimplemented addressing mode for instruction: %s", encodedInstruction.assembly)
	}

	return append([]byte{encodedInstruction.instruction.Opcode}, operandBytes...), nil
}
