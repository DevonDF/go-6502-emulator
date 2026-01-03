package assembler

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type Assembler struct {
	loadAddress  uint16                        // the address of where the program will be loaded within memory.
	bytecode     []byte                        // the raw bytecode generated for this program.
	instructions []EncodedInstructionInterface // a list of encoded instructions for the assembly program.
	labels       map[string]uint16             // a map of label names to relative address within the program.
	variables    map[string]string             // a map of variable names to the raw operand.
}

type EncodedInstructionInterface interface {
	ToByteCode(assembler *Assembler) ([]byte, error)
}

// NewAssembler creates and returns a new Assembler for use.
func NewAssembler(loadAddress uint16) *Assembler {
	return &Assembler{
		loadAddress:  loadAddress,
		bytecode:     make([]byte, 0),
		instructions: make([]EncodedInstructionInterface, 0),
		labels:       map[string]uint16{},
		variables:    map[string]string{},
	}
}

// getAddressingMode returns the AddressingMode used in the provided operands string and whether a label was used.
func (assembler *Assembler) getAddressingMode(operands string) (instructions.AddressingMode, bool, error) {
	var addressingMode instructions.AddressingMode
	labelUsed := false

	if operands == "" { // implied addressing, i.e. no operands
		addressingMode = instructions.AddrImplied
	} else if operands == "A" { // accumulator addressing
		addressingMode = instructions.AddrAccumulator
	} else if operands[0] == '#' { // immediate addressing
		addressingMode = instructions.AddrImmediate
	} else if operands[0] == '$' { // zeropage/absolute
		beforeComma, afterComma, commaFound := strings.Cut(operands, ",")
		if !commaFound {
			if len(operands) == 3 { // zeropage addressing
				addressingMode = instructions.AddrZeropage
			} else { // absolute addressing
				addressingMode = instructions.AddrAbsolute
			}
		} else { // X/Y zeropage or absolute addressing
			isZeroPage := len(beforeComma) == 3
			isAbsolute := !isZeroPage

			if isZeroPage && afterComma == "X" {
				addressingMode = instructions.AddrZeropageX
			} else if isZeroPage && afterComma == "Y" {
				addressingMode = instructions.AddrAbsoluteY
			} else if isAbsolute && afterComma == "X" {
				addressingMode = instructions.AddrAbsoluteX
			} else if isAbsolute && afterComma == "Y" {
				addressingMode = instructions.AddrAbsoluteY
			} else {
				return 0x0, labelUsed, fmt.Errorf("unrecognised addressing mode: %s", operands)
			}
		}
	} else if operands[0] == '(' { // indirect addressing
		if operands[len(operands)-1] == ')' {
			// indirectX addressing
			addressingMode = instructions.AddrIndirectX
		} else if operands[len(operands)-1] == 'Y' {
			addressingMode = instructions.AddrIndirectY
		} else {
			return 0x0, labelUsed, fmt.Errorf("unimplemented use of addressing mode: %s", operands)
		}
	} else {
		// assume this is a label
		addressingMode = instructions.AddrImplied
		labelUsed = true
	}

	return addressingMode, labelUsed, nil
}

// getInstructionFromAssemblyLine returns a pointer to an EncodedInstruction for this line of assembly.
func (assembler *Assembler) getInstructionFromAssemblyLine(line string) (*EncodedInstruction, error) {
	// split out instruction and operands
	assemblyString, operands, found := strings.Cut(line, " ")
	if !found {
		// command with no operand?
		instruction := instructions.InstructionFromAssembly(assemblyString, instructions.AddrImplied)
		if instruction != nil {
			return &EncodedInstruction{
				assembly:    line,
				operands:    "",
				instruction: instruction,
			}, nil
		}

		return nil, fmt.Errorf("invalid assembly line encountered: %s", line)
	}

	// check if the operand is a variable
	// as variables need to be defined before being used, we can be certain at this stage
	// we need to check first if some assembler-arithmatic is being used upon the variable
	beforePlus, afterPlus, foundPlus := strings.Cut(operands, "+")
	beforeMinus, afterMinus, foundMinus := strings.Cut(operands, "-")

	// find the actual variable if it exists
	foundVariable := false
	variableOperandPlus, foundBeforePlusVar := assembler.variables[beforePlus]
	if foundBeforePlusVar {
		operands = variableOperandPlus
		foundVariable = true
	}
	variableOperandMinus, foundBeforeMinusVar := assembler.variables[beforeMinus]
	if foundBeforeMinusVar {
		operands = variableOperandMinus
		foundVariable = true
	}

	// discover the addressing mode of the operand
	addressingMode, labelUsed, err := assembler.getAddressingMode(operands)
	if err != nil {
		return nil, err
	}

	// now we have the addressing mode, we can perform the assembler-arithmatic
	if foundVariable && (foundPlus || foundMinus) {
		var varValue uint64
		var format string

		switch addressingMode {
		case instructions.AddrImmediate:
			varValue, err = strconv.ParseUint(operands[2:], 16, 8)
			format = "#$%02X"
			if err != nil {
				return nil, fmt.Errorf("failed to parse uint from operand in instruction %s: %v", assemblyString, err)
			}
		case instructions.AddrZeropage:
			varValue, err = strconv.ParseUint(operands[1:], 16, 8)
			format = "$%02X"
			if err != nil {
				return nil, fmt.Errorf("failed to parse uint from operand in instruction %s: %v", assemblyString, err)
			}
		case instructions.AddrAbsolute:
			varValue, err = strconv.ParseUint(operands[1:], 16, 16)
			format = "$%04X"
			if err != nil {
				return nil, fmt.Errorf("failed to parse uint from operand in instruction %s: %v", assemblyString, err)
			}
		default:
			return nil, fmt.Errorf("attempted assembler-arithmatic on non-supported addressing type in instruction %s %s", assemblyString, operands)
		}

		if foundPlus {
			plusVal, err := strconv.ParseUint(afterPlus, 10, 8)
			if err != nil {
				return nil, fmt.Errorf("failed to parse uint from operand in instruction %s: %v", assemblyString, err)
			}
			operands = fmt.Sprintf(format, varValue+plusVal)
		} else {
			minusVal, err := strconv.ParseUint(afterMinus, 10, 8)
			if err != nil {
				return nil, fmt.Errorf("failed to parse uint from operand in instruction %s: %v", assemblyString, err)
			}
			operands = fmt.Sprintf(format, varValue-minusVal)
		}

	}

	var instruction *instructions.Instruction

	if !labelUsed || addressingMode == instructions.AddrIndirect {
		instruction = instructions.InstructionFromAssembly(assemblyString, addressingMode)
	} else {
		instruction = instructions.InstructionFromAssembly(assemblyString, instructions.AddrRelative)
		if instruction == nil {
			instruction = instructions.InstructionFromAssembly(assemblyString, instructions.AddrAbsolute)
		}
	}

	if instruction == nil {
		return nil, fmt.Errorf("unknown instruction: %s", line)
	}

	return &EncodedInstruction{
		assembly:       line,
		operands:       operands,
		instruction:    instruction,
		operandIsLabel: labelUsed,
	}, nil
}

// getAddressForLabel returns the address for a given label.
func (assembler *Assembler) getAddressForLabel(label string) (uint16, bool) {
	addr, ok := assembler.labels[label]
	return addr, ok
}

// Assemble takes a path to an assembly file and assembles it into bytecode
func (assembler *Assembler) Assemble(inputFilePath string) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// first pass - deliminate instructions by lines and assign
	// relative addresses
	// also handle labels here
	scanner := bufio.NewScanner(inputFile)
	relativeAddr := assembler.loadAddress
	for scanner.Scan() {
		line := scanner.Text()

		// sanitise comments
		beforeComment, _, commentExists := strings.Cut(line, ";")
		if commentExists {
			line = beforeComment
		}
		line = strings.TrimSpace(line)

		// skip if empty line
		if line == "" {
			continue
		}

		// check for defining assembler variables here
		if strings.HasPrefix(line, "define") {
			definitionSplit := strings.Split(line, " ")
			if len(definitionSplit) != 3 {
				return fmt.Errorf("incorrect usage of 'define': %s", line)
			}
			variableName := definitionSplit[1]
			operandValue := definitionSplit[2]

			_, found := assembler.labels[variableName]
			if found {
				return fmt.Errorf("conflicting variable name %s with label", variableName)
			}

			assembler.variables[variableName] = operandValue
			continue
		}

		// check for a label
		if strings.HasSuffix(line, ":") {
			labelName := line[:len(line)-1]

			_, found := assembler.variables[labelName]
			if found {
				return fmt.Errorf("conflicting label name %s with variable", labelName)
			}

			assembler.labels[labelName] = relativeAddr
			continue
		}

		// Create and append an EncodedInstructionInterface

		var encodedInstruction EncodedInstructionInterface
		var encodedInstructionSize uint16

		if strings.HasPrefix(line, ".byte") {
			// this is defining some raw byte memory
			encodedInstruction = NewEncodedMemoryInstruction(line, relativeAddr)
			bytecode, err := encodedInstruction.ToByteCode(assembler)
			if err != nil {
				return err
			}
			encodedInstructionSize = uint16(len(bytecode))
		} else {
			// this is an instruction
			inst, err := assembler.getInstructionFromAssemblyLine(line)
			if err != nil {
				return err
			}
			inst.relativeAddress = relativeAddr
			encodedInstructionSize = uint16(inst.instruction.Size)
			encodedInstruction = inst
		}

		assembler.instructions = append(assembler.instructions, encodedInstruction)
		relativeAddr += encodedInstructionSize
	}

	// now on second pass we can start compiling into bytecode
	for _, encodedInst := range assembler.instructions {
		instBytecode, err := encodedInst.ToByteCode(assembler)
		if err != nil {
			return err
		}
		assembler.bytecode = append(assembler.bytecode, instBytecode...)
	}

	return nil
}

// ToHexDump returns a hexdump of the bytecode generated
func (assembler *Assembler) ToHexDump() string {
	return hex.Dump(assembler.bytecode)
}

// Write writes any assembled bytecode to the provided file path.
func (assembler *Assembler) Write(outputFilePath string) error {
	err := os.WriteFile(outputFilePath, assembler.bytecode, 0644)
	if err != nil {
		return err
	}
	return nil
}
