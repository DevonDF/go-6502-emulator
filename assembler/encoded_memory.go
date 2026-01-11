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

func (encodedMemory *AssemblyInstructionMemory) decodeBytesTokens() ([]byte, error) {
	bytesTokens := encodedMemory.tokens[2:]
	bytecode := make([]byte, 0)

	for tokenIndex := 0; tokenIndex < len(bytesTokens); tokenIndex++ {
		token := bytesTokens[tokenIndex]
		if token.Type == TokenTypeSymbol && token.Token == "\"" {
			// read a string
			for _, strToken := range bytesTokens[tokenIndex+1:] {
				if strToken.Type == TokenTypeSymbol && strToken.Token == "\"" {
					break
				} else {
					for _, char := range strToken.Token {
						bytecode = append(bytecode, petscii.ASCIIToPetcsii[byte(char)])
					}
				}
				tokenIndex++
			}
		} else if token.Type == TokenTypeNumber {
			val, err := strconv.ParseUint(token.Token, 16, 8)
			if err != nil {
				return nil, fmt.Errorf("invalid number %s: %v", token.Token, err)
			}
			bytecode = append(bytecode, byte(val))
		} else if token.Type == TokenTypeString {
			// may be a hex string like 0x00
			if strings.HasPrefix(token.Token, "0x") {
				val, err := strconv.ParseUint(token.Token[2:], 16, 8)
				if err != nil {
					return nil, fmt.Errorf("invalid number %s: %v", token.Token, err)
				}
				bytecode = append(bytecode, byte(val))
			}
		}
		tokenIndex++
	}

	return bytecode, nil
}

func (encodedMemory *AssemblyInstructionMemory) ToByteCode(assembler *Assembler) ([]byte, error) {
	memoryTypeToken := encodedMemory.tokens[1]

	switch memoryTypeToken.Token {

	case "byte":
		return encodedMemory.decodeBytesTokens()
	default:
		return nil, fmt.Errorf("invalid memory type %s", memoryTypeToken.Token)

	}
}
