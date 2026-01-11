package assembler

import (
	"fmt"
	"strconv"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type AssemblyInstructionInstruction struct {
	instruction *instructions.Instruction // the assembly instruction this relates to.
	RawAssemblyInstruction
}

func (inst *AssemblyInstructionInstruction) ToByteCode(assembler *Assembler) ([]byte, error) {

	// Immediate or accumulator instructions are just single-byte opcodes
	if inst.instruction != nil && (inst.instruction.AddressingMode == instructions.AddrImplied ||
		inst.instruction.AddressingMode == instructions.AddrAccumulator) {
		return []byte{inst.instruction.Opcode}, nil
	}

	// Do a first-parse where we replace any definitions & labels
	// and perform any assembly-level arithmetic or operations
	opcodeToken := inst.tokens[0]
	operandTokens := inst.tokens[1:]
	parsedOperandTokens := make([]Token, 0)

	for tokenIndex, token := range operandTokens {
		if token.Type == TokenTypeString {
			// check if this is a label or definition
			labelInst, isLabel := assembler.getLabelByLabelName(token.Token)
			definition, isDefinition := assembler.definitions[token.Token]

			// assembly-operators
			isLowByte := false
			isHighByte := false
			toAdd := 0

			// handle any preceding operator
			if tokenIndex > 0 {
				previousToken := operandTokens[tokenIndex-1]
				if previousToken.Type == TokenTypeSymbol {
					switch previousToken.Token {
					case "<":
						isLowByte = true
					case ">":
						isHighByte = true
					}
				}
			}

			// handle any post operator
			if len(operandTokens) > tokenIndex+1 {
				nextToken := operandTokens[tokenIndex+1]
				if nextToken.Type == TokenTypeSymbol && (nextToken.Token == "+" || nextToken.Token == "-") {
					if len(operandTokens) < tokenIndex+2 {
						return nil, fmt.Errorf("invalid arithmetic - expected number after %s", nextToken.Token)
					}
					numToken := operandTokens[tokenIndex+2]
					numValue, err := strconv.Atoi(numToken.Token)
					if err != nil {
						return nil, fmt.Errorf("invalid number %s after %s", numToken.Token, nextToken.Token)
					}
					switch nextToken.Token {
					case "+":
						toAdd = numValue
					case "-":
						toAdd = -numValue
					}
					tokenIndex += 2
				}
			}

			if isLabel {
				// emit new tokens for the raw absolute address
				parsedOperandTokens = append(parsedOperandTokens,
					Token{Token: "$", Type: TokenTypeSymbol})

				addrToken := Token{Token: "", Type: TokenTypeNumber}
				if isLowByte {
					addrToken.Token = fmt.Sprintf("%02X", labelInst.relativeAddr&0xFF)
				} else if isHighByte {
					addrToken.Token = fmt.Sprintf("%02X", (labelInst.relativeAddr>>8)&0xFF)
				} else {
					addrToken.Token = fmt.Sprintf("%04X", labelInst.relativeAddr+uint16(toAdd))
				}

				parsedOperandTokens = append(parsedOperandTokens, addrToken)
			} else if isDefinition {
				// emit new tokens from the definition tokens
				parsedOperandTokens = append(parsedOperandTokens, definition.tokens[2:len(definition.tokens)-1]...)
				definitionNumToken := definition.tokens[len(definition.tokens)-1]
				numValue, err := strconv.Atoi(definitionNumToken.Token)
				if err != nil {
					return nil, fmt.Errorf("invalid definition number %s", definitionNumToken.Token)
				}

				if isLowByte {
					numValue = numValue & 0xFF
				} else if isHighByte {
					numValue = (numValue >> 8) & 0xFF
				} else {
					numValue = numValue + toAdd
				}

				numberToken := Token{Token: "", Type: TokenTypeNumber}

				switch definition.addressingMode {
				case instructions.AddrImmediate, instructions.AddrZeropage:
					numberToken.Token = fmt.Sprintf("%02X", numValue)
				case instructions.AddrAbsolute:
					numberToken.Token = fmt.Sprintf("%04X", numValue)
				}

				parsedOperandTokens = append(parsedOperandTokens, numberToken)

			} else {
				parsedOperandTokens = append(parsedOperandTokens, token)
			}
		} else if token.Type == TokenTypeSymbol && (token.Token == "<" || token.Token == ">") {
			continue
		} else {
			parsedOperandTokens = append(parsedOperandTokens, token)
		}
	}

	// if we don't already have the instruction, now we've pre-parsed the operand, we can get it
	if inst.instruction == nil {
		// now we have pre-parsed the operand, we can get the addressing mode for this instruction
		addrMode, err := assembler.getAddressingModeFromOperand(parsedOperandTokens)
		if err != nil {
			return nil, fmt.Errorf("failed to find addressing mode for operand: %v", err)
		}
		inst.instruction = instructions.InstructionFromAssembly(opcodeToken.Token, addrMode)

		if inst.instruction == nil {
			return nil, fmt.Errorf("failed to find instruction for line %s", inst.rawAssemblyLine)
		}
	}

	var operandBytes []byte
	var err error

	switch inst.instruction.AddressingMode {
	case instructions.AddrImplied:
		operandBytes = []byte{}

	case instructions.AddrAccumulator:
		operandBytes = []byte{}

	case instructions.AddrRelative:
		labelAddr, _ := strconv.ParseUint(parsedOperandTokens[0].Token, 16, 16)

		operandBytes = []byte{byte((uint16(labelAddr) - inst.relativeAddr) - uint16(inst.instruction.Size))}

	case instructions.AddrImmediate:
		operandBytes, err = assembler.hexStringToBytes(parsedOperandTokens[2].Token)

	case instructions.AddrZeropage, instructions.AddrZeropageX, instructions.AddrZeropageY:
		operandBytes, err = assembler.hexStringToBytes(parsedOperandTokens[1].Token)

	case instructions.AddrAbsolute, instructions.AddrAbsoluteX, instructions.AddrAbsoluteY:
		operandBytes, err = assembler.hexStringToBytes(parsedOperandTokens[1].Token)

	case instructions.AddrIndirectX, instructions.AddrIndirectY:
		operandBytes, err = assembler.hexStringToBytes(parsedOperandTokens[2].Token)

	default:
		return nil, fmt.Errorf("unimplemented addressing mode for instruction: %s", inst.rawAssemblyLine)
	}

	if err != nil {
		return nil, err
	}

	return append([]byte{inst.instruction.Opcode}, operandBytes...), nil
}
