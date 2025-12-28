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
			instruction: &Instruction{
				opcode:  0x5,
				size:    2,
				cycles:  3,
				handler: nil,
			},
			operands: []byte{0x01},
		},
		inst)
}

func TestDecodeNextInstruction_NonInstruction(test *testing.T) {
	data := []byte{0xFF, 0x15, 0xF2}
	reader := bufio.NewReader(bytes.NewReader(data))

	_, err := DecodeNextInstruction(reader)

	assert.Error(test, err)
}
