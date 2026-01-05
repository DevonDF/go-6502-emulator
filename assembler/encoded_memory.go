package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DevonDF/go-6502-emulator/emulator/petscii"
)

type AssemblyInstructionMemory struct {
	RawAssemblyInstruction
}

func (encodedMemory *AssemblyInstructionMemory) ToByteCode(assembler *Assembler) ([]byte, error) {
	operands := strings.TrimPrefix(encodedMemory.rawAssemblyLine, ".byte ")
	operandsSplit := strings.Split(operands, ",")

	bytecode := make([]byte, 0)

	for _, operand := range operandsSplit {
		operand = strings.TrimSpace(operand)

		if operand[0] == '"' && operand[len(operand)-1] == '"' {
			// string
			for _, char := range operand[1 : len(operand)-1] {
				bytecode = append(bytecode, petscii.ASCIIToPetcsii[byte(char)])
			}
		} else if strings.HasPrefix(operand, "0x") {
			// hex
			val, err := strconv.ParseUint(operand[2:], 16, 8)
			if err != nil {
				return nil, fmt.Errorf("invalid operand %s: %v", operand, err)
			}
			bytecode = append(bytecode, byte(val))
		}

	}

	return bytecode, nil
}
