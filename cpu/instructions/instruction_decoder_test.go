package instructions

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNextInstruction_CorrectInstruction(test *testing.T) {
	data := []byte{0x05, 0x01}
	reader := bufio.NewReader(bytes.NewReader(data))

	inst, err := DecodeNextInstruction(reader)

	assert.NoError(test, err)

	assert.Equal(test,
		DecodedInstruction{
			Instruction: &Instruction{
				Opcode:  0x5,
				Size:    2,
				Cycles:  3,
				Handler: nil,
			},
			Operands: []byte{0x01},
		},
		inst)
}

func TestDecodeNextInstruction_NonInstruction(test *testing.T) {
	data := []byte{0xFF, 0x15, 0xF2}
	reader := bufio.NewReader(bytes.NewReader(data))

	_, err := DecodeNextInstruction(reader)

	assert.Error(test, err)
}
