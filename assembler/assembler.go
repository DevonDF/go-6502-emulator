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
	instructions []EncodedInstructionInterface // a list of instructions for the assembly program.
	definitions  map[string]AssemblyDefinition // a map of definitions names to the AssemblyDefinition struct.
	labels       map[string]bool               // a map of declared labels.
}

// RawAssemblyInstruction is a template struct for an assembly instruction.
type RawAssemblyInstruction struct {
	rawAssemblyLine string // the raw assembly line.
	memorySize      uint16 // the size of this instruction in raw bytes.
	relativeAddr    uint16 // the relative address.
}

type EncodedInstructionInterface interface {
	ToByteCode(assembler *Assembler) ([]byte, error)
}

// AssemblyDefinition represents a defined variable within assembly.
type AssemblyDefinition struct {
	name           string                      // the name of the variable
	rawOperand     string                      // the raw string operand for replacement
	intOperand     uint16                      // the integer representation of the operand
	addressingMode instructions.AddressingMode // the addressing mode of the operand
}

// NewAssembler creates and returns a new Assembler for use.
func NewAssembler(loadAddress uint16) *Assembler {
	return &Assembler{
		loadAddress:  loadAddress,
		bytecode:     make([]byte, 0),
		instructions: make([]EncodedInstructionInterface, 0),
		definitions:  map[string]AssemblyDefinition{},
		labels:       map[string]bool{},
	}
}

// stringContainsLabel returns whether a string contains a label defined within the assembler.
func (assembler *Assembler) stringContainsLabel(str string) bool {
	for label, _ := range assembler.labels {
		if strings.Contains(str, label) {
			return true
		}
	}
	return false
}

func (assembler *Assembler) getLabelWithinOperand(operand string) (*AssemblyInstructionLabel, bool) {
	for _, assemblyLine := range assembler.instructions {
		label, ok := assemblyLine.(*AssemblyInstructionLabel)
		if !ok {
			continue
		}
		if strings.Contains(operand, label.labelName) {
			return label, true
		}
	}
	return nil, false
}

// getDefinitionWithinOperands returns whether a string contains a definition defined within the assembler.
func (assembler *Assembler) getDefinitionWithinOperand(operand string) (*AssemblyDefinition, bool) {
	for varName, definition := range assembler.definitions {
		if strings.Contains(operand, varName) {
			return &definition, true
		}
	}
	return nil, false
}

func (assembler *Assembler) readNumberStringFromString(str string) string {
	intStr := ""
	for _, chRune := range str {
		char := string(chRune)
		_, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		intStr += char
	}
	return intStr
}

// preParseAssemblyOperand performs a pre-parse by replacing any labels & definitions and handling any math.
func (assembler *Assembler) preParseAssemblyOperand(operand string) (string, error) {
	newOperand := operand

	// handle assembler definitions
	definition, hasDefinition := assembler.getDefinitionWithinOperand(operand)
	if hasDefinition {
		value := definition.intOperand

		// handle any arithmetic
		definitionStartIndex := strings.Index(operand, definition.name)
		definitionEndIndex := definitionStartIndex + len(definition.name)
		if len(operand) > definitionEndIndex {
			charAfter := operand[definitionEndIndex]
			switch charAfter {
			case '+':
				numStr := assembler.readNumberStringFromString(operand[definitionEndIndex+1:])
				toAdd, _ := strconv.Atoi(numStr)
				value += uint16(toAdd)
				definitionEndIndex += len(numStr) + 1
			case '-':
				numStr := assembler.readNumberStringFromString(operand[definitionEndIndex+1:])
				toMinus, _ := strconv.Atoi(numStr)
				value += uint16(toMinus)
				definitionEndIndex += len(numStr) + 1
			}
		}

		// create the string for this new value
		valueStr := ""
		switch definition.addressingMode {
		case instructions.AddrImmediate:
			valueStr = fmt.Sprintf("#$%02X", value)
		case instructions.AddrZeropage:
			valueStr = fmt.Sprintf("$%02X", value)
		case instructions.AddrAbsolute:
			valueStr = fmt.Sprintf("$%04X", value)
		}

		// write back the new operand
		newOperand = newOperand[:definitionStartIndex] + valueStr + newOperand[definitionEndIndex:]
	}

	// handle labels
	label, hasLabel := assembler.getLabelWithinOperand(operand)
	if hasLabel {
		labelAddr := label.relativeAddr

		// handle any arithmetic
		labelStartIndex := strings.Index(operand, label.labelName)
		labelEndIndex := labelStartIndex + len(label.labelName)
		if len(operand) > labelEndIndex {
			charAfter := operand[labelEndIndex]
			switch charAfter {
			case '+':
				numStr := assembler.readNumberStringFromString(operand[labelEndIndex+1:])
				toAdd, _ := strconv.Atoi(numStr)
				labelAddr += uint16(toAdd)
				labelEndIndex += len(numStr) + 1
			case '-':
				numStr := assembler.readNumberStringFromString(operand[labelEndIndex+1:])
				toMinus, _ := strconv.Atoi(numStr)
				labelAddr += uint16(toMinus)
				labelEndIndex += len(numStr) + 1
			}
		}

		// labels are always absolute addresses
		// write back the new operand
		newOperand = newOperand[:labelStartIndex] + fmt.Sprintf("$%04X", labelAddr) + newOperand[labelEndIndex:]
	}

	return newOperand, nil
}

// getAddressingMode returns the AddressingMode used in the provided operands string and whether a label was used.
func (assembler *Assembler) getAddressingMode(operands string) (instructions.AddressingMode, error) {
	var addressingMode instructions.AddressingMode

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
				return 0x0, fmt.Errorf("unrecognised addressing mode: %s", operands)
			}
		}
	} else if operands[0] == '(' { // indirect addressing
		if operands[len(operands)-1] == ')' {
			// indirectX addressing
			addressingMode = instructions.AddrIndirectX
		} else if operands[len(operands)-1] == 'Y' {
			addressingMode = instructions.AddrIndirectY
		} else {
			return 0x0, fmt.Errorf("unimplemented use of addressing mode: %s", operands)
		}
	} else {
		// assume this is a label
		addressingMode = instructions.AddrImplied
	}

	return addressingMode, nil
}

// hexStringToBytes parses a hex string and returns a little-endian encoded byte array.
func (assembler *Assembler) hexStringToBytes(hexStr string) ([]byte, error) {
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

// Assemble takes a path to an assembly file and assembles it into bytecode
func (assembler *Assembler) Assemble(inputFilePath string) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// first pass - convert everything to an AssemblyInstructions & handle definitions
	scanner := bufio.NewScanner(inputFile)
	assembler.instructions = make([]EncodedInstructionInterface, 0)
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

			var addressingMode instructions.AddressingMode
			var intOperand uint16

			if strings.HasPrefix(operandValue, "#$") {
				// 1byte immediate addressing
				addressingMode = instructions.AddrImmediate
				parsedInt, err := strconv.ParseUint(operandValue[2:], 16, 8)
				if err != nil {
					return fmt.Errorf("failed to parse integer in definition '%s': %v", line, err)
				}
				intOperand = uint16(parsedInt)
			} else if strings.HasPrefix(operandValue, "$") && len(operandValue) == 3 {
				// 1byte zeropage addressing
				addressingMode = instructions.AddrZeropage
				parsedInt, err := strconv.ParseUint(operandValue[1:], 16, 8)
				if err != nil {
					return fmt.Errorf("failed to parse integer in definition '%s': %v", line, err)
				}
				intOperand = uint16(parsedInt)
			} else if strings.HasPrefix(operandValue, "$") && len(operandValue) == 5 {
				// 2byte absolute addressing
				addressingMode = instructions.AddrAbsolute
				parsedInt, err := strconv.ParseUint(operandValue[1:], 16, 16)
				if err != nil {
					return fmt.Errorf("failed to parse integer in definition '%s': %v", line, err)
				}
				intOperand = uint16(parsedInt)
			} else {
				return fmt.Errorf("failed to parse definition '%s': no valid addressing found", line)
			}

			assembler.definitions[variableName] = AssemblyDefinition{
				name:           variableName,
				rawOperand:     operandValue,
				addressingMode: addressingMode,
				intOperand:     intOperand,
			}
			continue
		}

		// Handle labels & instructions
		if strings.HasSuffix(line, ":") {
			// this is a line of a label
			labelName := line[:len(line)-1]

			_, found := assembler.definitions[labelName]
			if found {
				return fmt.Errorf("conflicting label name %s with variable", labelName)
			}

			assembler.instructions = append(assembler.instructions, &AssemblyInstructionLabel{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
				},
				labelName: labelName,
			})
			assembler.labels[labelName] = true
		} else if strings.HasPrefix(line, ".byte") {
			// this is a line defining some raw byte memory
			assembler.instructions = append(assembler.instructions, &AssemblyInstructionMemory{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
				},
			})
		} else {
			// this is an assembly instruction
			assembler.instructions = append(assembler.instructions, &AssemblyInstructionInstruction{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
				},
			})
		}
	}

	// Now we have sanitised the assembly code into a simple list of AssemblyCode structs
	// Here we can now assign a memorySize to each AssemblyCodeLine in order to
	relativeAddr := assembler.loadAddress
	for _, assemblyLine := range assembler.instructions {
		switch inst := assemblyLine.(type) {

		case *AssemblyInstructionLabel:
			inst.memorySize = uint16(0)
			inst.relativeAddr = relativeAddr

		case *AssemblyInstructionMemory:
			// here we can compile the memory declaration and get the size manually
			bytecode, err := inst.ToByteCode(assembler)
			if err != nil {
				return fmt.Errorf("failed to generate bytecode for memory declaration %s: %v", inst.rawAssemblyLine, err)
			}
			inst.memorySize = uint16(len(bytecode))
			inst.relativeAddr = relativeAddr
			relativeAddr += inst.memorySize

		case *AssemblyInstructionInstruction:
			// here we need to get the addressing mode used for this instruction
			// and then find the instruction
			var instruction *instructions.Instruction

			opcode, operand, hasOperand := strings.Cut(inst.rawAssemblyLine, " ")
			inst.opcodeString = opcode
			inst.operand = operand

			if !hasOperand {
				instruction = instructions.InstructionFromAssembly(inst.rawAssemblyLine, instructions.AddrImplied)
				// this is the actual instruction, so may aswell keep it here
				inst.instruction = instruction
			} else {
				definition, foundDefinition := assembler.getDefinitionWithinOperand(operand)
				_, foundLabel := assembler.getLabelWithinOperand(operand)
				if foundDefinition {
					// if it contains a definition, grab that to understand the addressing mode here
					instruction = instructions.InstructionFromAssembly(opcode, definition.addressingMode)
				} else if foundLabel {
					instruction = instructions.InstructionFromAssembly(opcode, instructions.AddrRelative)
					if instruction != nil {
						// this is the actual instruction, so may aswell keep it here
						inst.instruction = instruction
					} else { // assume just absolute at this point, but it could be absolute,X or absolute,Y
						instruction = instructions.InstructionFromAssembly(opcode, instructions.AddrAbsolute)
					}
				} else {
					addrMode, err := assembler.getAddressingMode(operand)
					if err != nil {
						return fmt.Errorf("invalid operand in line %s", inst.rawAssemblyLine)
					}
					instruction = instructions.InstructionFromAssembly(opcode, addrMode)
					// this is the actual instruction, so may aswell keep it here
					inst.instruction = instruction
				}
			}

			if instruction == nil {
				return fmt.Errorf("unknown instruction %s", inst.rawAssemblyLine)
			}

			inst.memorySize = uint16(instruction.Size)
			inst.relativeAddr = relativeAddr
			relativeAddr += inst.memorySize
		}
	}

	// Now our list of AbstractAssemblyInstructions should all have valid memory sizes and relative addresses
	// We will do our final pass to assemble each line into bytecode
	assembler.bytecode = make([]byte, 0)

	for _, assemblyLine := range assembler.instructions {
		switch inst := assemblyLine.(type) {

		case *AssemblyInstructionMemory:
			bytecode, err := inst.ToByteCode(assembler)
			if err != nil {
				return fmt.Errorf("failed to assemble instruction %s: %v", inst.rawAssemblyLine, err)
			}
			assembler.bytecode = append(assembler.bytecode, bytecode...)

		case *AssemblyInstructionInstruction:
			bytecode, err := inst.ToByteCode(assembler)
			if err != nil {
				return fmt.Errorf("failed to assemble instruction %s: %v", inst.rawAssemblyLine, err)
			}
			assembler.bytecode = append(assembler.bytecode, bytecode...)
		}
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
