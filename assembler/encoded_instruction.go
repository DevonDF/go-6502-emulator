package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type AssemblyInstructionInstruction struct {
	opcodeString string                    // the opcode string, e.g. ADC.
	operand      string                    // the string operand.
	instruction  *instructions.Instruction // the assembly instruction this relates to.
	RawAssemblyInstruction
}

func (inst *AssemblyInstructionInstruction) ToByteCode(assembler *Assembler) ([]byte, error) {
	if inst.operand == "" {
		// implied addressing, find the instruction & return the opcode as a byte
		instruction := instructions.InstructionFromAssembly(inst.opcodeString, instructions.AddrImplied)
		if instruction == nil {
			return nil, fmt.Errorf("failed to find instruction for opcode %s", inst.opcodeString)
		}
		return []byte{instruction.Opcode}, nil
	}

	newOperand, err := assembler.preParseAssemblyOperand(inst.operand)
	if err != nil {
		return nil, err
	}

	// if we don't already have the instruction, now we've pre-parsed the operand, we can get it
	if inst.instruction == nil {
		// now we have pre-parsed the operand, we can get the addressing mode for this instruction
		addrMode, err := assembler.getAddressingMode(newOperand)
		if err != nil {
			return nil, fmt.Errorf("failed to find addressing mode for operand %s: %v", newOperand, err)
		}
		inst.instruction = instructions.InstructionFromAssembly(inst.opcodeString, addrMode)

		if inst.instruction == nil {
			return nil, fmt.Errorf("failed to find instruction for line %s", inst.rawAssemblyLine)
		}
	}

	var operandBytes []byte

	switch inst.instruction.AddressingMode {
	case instructions.AddrImplied:
		operandBytes = []byte{}

	case instructions.AddrAccumulator:
		operandBytes = []byte{}

	case instructions.AddrRelative:
		labelAddr, _ := strconv.ParseUint(newOperand[2:], 16, 16)

		operandBytes = []byte{byte((uint16(labelAddr) - inst.relativeAddr) - uint16(inst.instruction.Size))}

	case instructions.AddrImmediate:
		operandBytes, err = assembler.hexStringToBytes(newOperand[2:])

	case instructions.AddrZeropage:
		operandBytes, err = assembler.hexStringToBytes(newOperand[1:])

	case instructions.AddrZeropageX, instructions.AddrZeropageY:
		beforeComma, _, _ := strings.Cut(newOperand, ",")
		operandBytes, err = assembler.hexStringToBytes(beforeComma[1:])

	case instructions.AddrAbsolute:
		operandBytes, err = assembler.hexStringToBytes(newOperand[1:])

	case instructions.AddrAbsoluteX, instructions.AddrAbsoluteY:
		beforeComma, _, _ := strings.Cut(newOperand, ",")
		operandBytes, err = assembler.hexStringToBytes(beforeComma[1:])

	case instructions.AddrIndirectX, instructions.AddrIndirectY:
		operandBytes, err = assembler.hexStringToBytes(newOperand[2:4])

	default:
		return nil, fmt.Errorf("unimplemented addressing mode for instruction: %s", inst.rawAssemblyLine)
	}

	if err != nil {
		return nil, err
	}

	return append([]byte{inst.instruction.Opcode}, operandBytes...), nil
}
