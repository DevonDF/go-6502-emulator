package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/emulator/petscii"
)

type EncodedMemoryInstruction struct {
	assembly                string // the raw assembly for this encoded memory.
	bytecode                []byte // the bytecode for this memory.
	errorGeneratingBytecode error  // any error generating the bytecode
	relativeAddress         uint16 // the relative address for this encoded memory.
}

func NewEncodedMemoryInstruction(assemblyLine string, relativeAddress uint16) *EncodedMemoryInstruction {
	mem := &EncodedMemoryInstruction{
		assembly:        assemblyLine,
		relativeAddress: relativeAddress,
	}
	mem.GenerateByteCode()
	return mem
}

// generateByteCode generates the bytecode for this memory instruction and populates the internal variable.
func (encodedMemory *EncodedMemoryInstruction) GenerateByteCode() {
	operands := strings.TrimPrefix(encodedMemory.assembly, ".byte ")
	operandsSplit := strings.Split(operands, ",")

	encodedMemory.bytecode = make([]byte, 0)

	for _, operand := range operandsSplit {
		operand = strings.TrimSpace(operand)

		if operand[0] == '"' && operand[len(operand)-1] == '"' {
			// string
			for _, char := range operand[1 : len(operand)-1] {
				encodedMemory.bytecode = append(encodedMemory.bytecode, petscii.ASCIIToPetcsii[byte(char)])
			}
		} else if strings.HasPrefix(operand, "0x") {
			// hex
			val, err := strconv.ParseUint(operand[2:], 16, 8)
			if err != nil {
				encodedMemory.errorGeneratingBytecode = fmt.Errorf("invalid operand %s: %v", operand, err)
				return
			}
			encodedMemory.bytecode = append(encodedMemory.bytecode, byte(val))
		}

	}
}

func (encodedMemory *EncodedMemoryInstruction) ToByteCode(assembler *Assembler) ([]byte, error) {
	return encodedMemory.bytecode, encodedMemory.errorGeneratingBytecode
}
