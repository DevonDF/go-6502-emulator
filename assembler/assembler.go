package assembler

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/DevonDF/go-6502-emulator/emulator/cpu/instructions"
)

type Assembler struct {
	loadAddress  uint16                              // the address of where the program will be loaded within memory.
	bytecode     []byte                              // the raw bytecode generated for this program.
	instructions []EncodedInstructionInterface       // a list of instructions for the assembly program.
	definitions  map[string]AssemblyDefinition       // a map of definitions names to the AssemblyDefinition struct.
	labels       map[string]AssemblyInstructionLabel // a map of declared labels.
}

// RawAssemblyInstruction is a template struct for an assembly instruction.
type RawAssemblyInstruction struct {
	rawAssemblyLine string  // the raw assembly line as a string.
	tokens          []Token // the tokens that make up this assembly instruction.
	memorySize      uint16  // the size of this instruction in raw bytes.
	relativeAddr    uint16  // the relative address.
}

type EncodedInstructionInterface interface {
	ToByteCode(assembler *Assembler) ([]byte, error)
}

// AssemblyDefinition represents a defined variable within assembly.
type AssemblyDefinition struct {
	name           string                      // the name of the variable
	tokens         []Token                     // the tokens for this definition
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
		labels:       map[string]AssemblyInstructionLabel{},
	}
}

func (assembler *Assembler) getLabelWithinOperand(operandTokens []Token) (*AssemblyInstructionLabel, bool) {
	for _, token := range operandTokens {
		for labelName, labelInst := range assembler.labels {
			if token.Token == labelName {
				return &labelInst, true
			}
		}
	}
	return nil, false
}

// getDefinitionWithinOperands returns whether a string contains a definition defined within the assembler.
func (assembler *Assembler) getDefinitionWithinOperand(operandTokens []Token) (*AssemblyDefinition, bool) {
	for _, token := range operandTokens {
		for varName, definition := range assembler.definitions {
			if token.Token == varName {
				return &definition, true
			}
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

// getAddressingModeFromOperand returns the AddressingMode used in the provided operand tokens.
func (assembler *Assembler) getAddressingModeFromOperand(operands []Token) (instructions.AddressingMode, error) {
	// Implied addressing
	if len(operands) == 0 {
		return instructions.AddrImplied, nil
	}

	// Accumulator addressing
	if operands[0].Type == TokenTypeString && operands[0].Token == "A" {
		return instructions.AddrAccumulator, nil
	}

	// Immediate addressing
	if operands[0].Type == TokenTypeSymbol && operands[0].Token == "#" {
		return instructions.AddrImmediate, nil
	}

	// Absolute | Absolute,X | Absolute,Y | Zeropage | Zeropage,X | Zeropage,Y addressing
	if operands[0].Type == TokenTypeSymbol && operands[0].Token == "$" {
		if len(operands) < 2 {
			return 0x0, fmt.Errorf("expected valid hex string after $")
		}
		hexNumToken := operands[1]
		if hexNumToken.Type != TokenTypeNumber {
			return 0x0, fmt.Errorf("expected valid hex string after $")
		}

		switch len(hexNumToken.Token) {
		case 2: // zeropage | zeropage,X | zeropage,Y
			if len(operands) == 2 {
				return instructions.AddrZeropage, nil
			} else if len(operands) == 4 {
				regToken := operands[3]
				switch regToken.Token {
				case "X":
					return instructions.AddrZeropageX, nil
				case "Y":
					return instructions.AddrZeropageY, nil
				default:
					return 0x0, fmt.Errorf("expected zeropage,X or zeropage,Y addressing")
				}
			} else {
				return 0x0, fmt.Errorf("invalid zeropage operand")
			}
		case 4: // absolute | absolute,X | absolute,Y
			if len(operands) == 2 {
				return instructions.AddrAbsolute, nil
			} else if len(operands) == 4 {
				regToken := operands[3]
				switch regToken.Token {
				case "X":
					return instructions.AddrAbsoluteX, nil
				case "Y":
					return instructions.AddrAbsoluteY, nil
				default:
					return 0x0, fmt.Errorf("expected absolute,X or absolute,Y addressing")
				}
			} else {
				return 0x0, fmt.Errorf("invalid absolute operand")
			}
		default:
			return 0x0, fmt.Errorf("expected absolute or zeropage addressing after $")
		}
	}

	// indirect addressing
	if operands[0].Type == TokenTypeSymbol && operands[0].Token == "(" {
		lastToken := operands[len(operands)-1]
		if lastToken.Type == TokenTypeSymbol && lastToken.Token == ")" {
			return instructions.AddrIndirectX, nil
		} else if lastToken.Type == TokenTypeString && lastToken.Token == "Y" {
			return instructions.AddrIndirectY, nil
		} else {
			return 0x0, fmt.Errorf("expected valid indirect addressing after (")
		}
	}

	return 0x0, fmt.Errorf("unrecognised operand addressing mode")
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

	// first pass - tokenise everything and convert to respective AssembleyInstruction
	scanner := bufio.NewScanner(inputFile)
	assembler.instructions = make([]EncodedInstructionInterface, 0)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := Tokenise(line)

		fmt.Printf("%s = %v\n", line, tokens)

		if len(tokens) == 0 {
			continue
		}

		if tokens[0].Type == TokenTypeString && tokens[0].Token == "define" {
			// handle definitions
			// define <string> $<number>
			// define <string> #$<number>
			if tokens[1].Type != TokenTypeString {
				return fmt.Errorf("incorrectly formatted definition %s: expected definition name but got %s",
					line, tokens[1].Token)
			}
			definitionName := tokens[1].Token

			newDefinition := AssemblyDefinition{
				name:   definitionName,
				tokens: tokens,
			}

			switch len(tokens) {
			case 4:
				if tokens[2].Type != TokenTypeSymbol || tokens[2].Token != "$" {
					return fmt.Errorf("incorrectly formatted definition %s: expected $ after %s", line, tokens[1].Token)
				}
				if tokens[3].Type != TokenTypeNumber {
					return fmt.Errorf("incorrectly formatted definition %s: expected number after %s", line, tokens[2].Token)
				}

				intVal, err := strconv.ParseUint(tokens[3].Token, 16, 0)
				if err != nil {
					return fmt.Errorf("incorrectly formatted definition %s: %s is not a valid number", line, tokens[3].Token)
				}
				newDefinition.intOperand = uint16(intVal)

				switch len(tokens[3].Token) {
				case 2:
					newDefinition.addressingMode = instructions.AddrZeropage
				case 4:
					newDefinition.addressingMode = instructions.AddrAbsolute
				default:
					return fmt.Errorf("incorrectly formatted definition %s: %s is not a valid zeropage or absolute address", line, tokens[3].Token)
				}
			case 5:
				if tokens[2].Type != TokenTypeSymbol || tokens[2].Token != "#" {
					return fmt.Errorf("incorrectly formatted definition %s: expected # after %s", line, tokens[1].Token)
				}
				if tokens[3].Type != TokenTypeSymbol || tokens[2].Token != "$" {
					return fmt.Errorf("incorrectly formatted definition %s: expected $ after %s", line, tokens[2].Token)
				}
				if tokens[4].Type != TokenTypeNumber {
					return fmt.Errorf("incorrectly formatted definition %s: expected number after %s", line, tokens[3].Token)
				}
				if len(tokens[4].Token) != 2 {
					return fmt.Errorf("incorrectly formatted definition %s: %s is not a valid hex-encoded immediate byte", line, tokens[4].Token)
				}

				intVal, err := strconv.ParseUint(tokens[3].Token, 16, 0)
				if err != nil {
					return fmt.Errorf("incorrectly formatted definition %s: %s is not a valid number", line, tokens[3].Token)
				}
				newDefinition.intOperand = uint16(intVal)
				newDefinition.addressingMode = instructions.AddrImmediate
			}

			assembler.definitions[definitionName] = newDefinition
			continue
		} else if tokens[len(tokens)-1].Type == TokenTypeSymbol && tokens[len(tokens)-1].Token == ":" {
			// Handle labels
			// <string>:
			if tokens[0].Type != TokenTypeString {
				return fmt.Errorf("incorrectly formatted line %s: expected label name before :", line)
			}
			if len(tokens) > 2 {
				return fmt.Errorf("incorrectly formatted label %s: expected <name>:", line)
			}
			labelName := tokens[0].Token

			_, found := assembler.definitions[labelName]
			if found {
				return fmt.Errorf("conflicting label name %s with variable", labelName)
			}

			instLabel := AssemblyInstructionLabel{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
					tokens:          tokens,
				},
				labelName: labelName,
			}

			assembler.instructions = append(assembler.instructions, &instLabel)
			assembler.labels[labelName] = instLabel
		} else if tokens[0].Type == TokenTypeSymbol && tokens[0].Token == "." {
			// Handle memory definitions
			// .<name> "<string>"
			// .<name> $<number>, $number,...
			// ...
			if tokens[1].Type != TokenTypeString {
				return fmt.Errorf("incorrectly formatted declaration of memory %s: expected a string after .", line)
			}
			assembler.instructions = append(assembler.instructions, &AssemblyInstructionMemory{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
					tokens:          tokens,
				},
			})
		} else if tokens[0].Type == TokenTypeString {
			// Handle instructions
			assembler.instructions = append(assembler.instructions, &AssemblyInstructionInstruction{
				RawAssemblyInstruction: RawAssemblyInstruction{
					rawAssemblyLine: line,
					tokens:          tokens,
				},
			})
		} else {
			return fmt.Errorf("unrecognised assembly line %s", line)
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
			// here we need to get the size of the instruction
			// we have enough knowledge in order to discover this now
			// but not the exact instruction in all cases
			var memorySize uint16

			opcodeToken := inst.tokens[0]
			hasOperand := len(inst.tokens) > 0

			if !hasOperand {
				// this is an non-operand implied opcode
				// find the instruction and set it
				instruction := instructions.InstructionFromAssembly(opcodeToken.Token, instructions.AddrImplied)
				if instruction == nil {
					return fmt.Errorf("invalid assembly code %s: %s not found or missing operand", inst.rawAssemblyLine, opcodeToken.Token)
				}
				inst.instruction = instruction
				memorySize = uint16(instruction.Size)
			} else {
				operandTokens := inst.tokens[1:]
				definition, hasDefinition := assembler.getDefinitionWithinOperand(operandTokens)
				_, foundLabel := assembler.getLabelWithinOperand(operandTokens)

				if hasDefinition {
					// if it contains a definition, we can use that computed addressing mode here
					switch definition.addressingMode {
					case instructions.AddrImmediate, instructions.AddrZeropage:
						memorySize = uint16(2)
					case instructions.AddrAbsolute:
						memorySize = uint16(3)
					}
				} else if foundLabel {
					// it contains a label, so it is either relative or absolute addressing
					instruction := instructions.InstructionFromAssembly(opcodeToken.Token, instructions.AddrRelative)
					if instruction != nil {
						// this is the actual instruction, so may aswell keep it here
						inst.instruction = instruction
						memorySize = uint16(instruction.Size)
					} else { // it is absolute addressing, so 3 bytes
						memorySize = uint16(3)
					}
				} else {
					addrMode, err := assembler.getAddressingModeFromOperand(operandTokens)
					if err != nil {
						return fmt.Errorf("invalid assembly code %s: invalid operand for opcode %s: %v", inst.rawAssemblyLine, opcodeToken.Token, err)
					}
					instruction := instructions.InstructionFromAssembly(opcodeToken.Token, addrMode)
					// this is the actual instruction, so may aswell keep it here
					inst.instruction = instruction
				}
			}

			inst.memorySize = memorySize
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
