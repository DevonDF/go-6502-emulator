package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type EncodedInstruction struct {
	assembly        string                    // the raw assembly for this instruction.
	operands        string                    // the operands string.
	instruction     *instructions.Instruction // the instruction of this instruction.
	relativeAddress int                       // the relative address of this instruction.
}

// hexStringToBytes parses a hex string and returns a little-endian encoded byte array.
func hexStringToBytes(hexStr string) ([]byte, error) {
	if len(hexStr) == 2 {
		v, err := strconv.ParseUint(hexStr, 16, 8)
		if err != nil {
			return nil, err
		}
		return []byte{byte(v)}, nil
	} else if len(hexStr) == 4 {
		v, err := strconv.ParseUint(hexStr, 16, 16)
		if err != nil {
			return nil, err
		}
		return []byte{byte(v & 0xFF), byte((v >> 8) & 0xFF)}, nil
	}
	return nil, fmt.Errorf("invalid hex string %s", hexStr)
}

// ToBytecode compiles and returns into bytecode the instruction.
func (encodedInstruction *EncodedInstruction) ToBytecode(assembler *Assembler) ([]byte, error) {
	var operandBytes []byte
	var err error

	switch encodedInstruction.instruction.AddressingMode {
	case instructions.AddrImplied:
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
		operandBytes, err = hexStringToBytes(encodedInstruction.operands[2:])

	case instructions.AddrZeropage:
		operandBytes, err = hexStringToBytes(encodedInstruction.operands[1:])

	case instructions.AddrZeropageX, instructions.AddrZeropageY:
		beforeComma, _, _ := strings.Cut(encodedInstruction.operands, ",")
		operandBytes, err = hexStringToBytes(beforeComma[1:])

	case instructions.AddrAbsolute:
		operandBytes, err = hexStringToBytes(encodedInstruction.operands[1:])

	case instructions.AddrAbsoluteX, instructions.AddrAbsoluteY:
		beforeComma, _, _ := strings.Cut(encodedInstruction.operands, ",")
		operandBytes, err = hexStringToBytes(beforeComma[1:])

	default:
		return nil, fmt.Errorf("unimplemented addressing mode for instruction: %s", encodedInstruction.assembly)
	}

	if err != nil {
		return nil, err
	}

	return append([]byte{encodedInstruction.instruction.Opcode}, operandBytes...), nil
}
