package memory

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"log/slog"
)

// $0000–$00FF  Zero Page
// $0100–$01FF  Stack
// $0200–$07FF  RAM
// $8000–$FFFF  ROM
const (
	ZeroPageAddress   = 0x0000
	StackStartAddress = 0x0100
	RAMStartAddress   = 0x0200
	ROMStartAddress   = 0x8000
)

type Memory struct {
	memory []byte
	logger *slog.Logger
}

// NewMemory creates and returns a new Memory object.
func NewMemory(logger *slog.Logger) *Memory {
	return &Memory{
		memory: make([]byte, 65536), // 6502 has maximum memory size of 16K due to 16bit PC
		logger: logger,
	}
}

// Write writes to the internal memory at the given offset the provided data in its entirety.
func (memory *Memory) Write(address uint16, data []byte) error {
	if int(address)+len(data) > len(memory.memory) {
		return errors.New("out of bound write")
	}
	copy(memory.memory[int(address):], data)
	memory.logger.Debug("wrote memory", "start", address, "end", int(address)+len(data), "size", len(data))
	return nil
}

// ReadByte reads a byte at a given memory address.
func (memory *Memory) ReadByte(address uint16) (byte, error) {
	if int(address) > len(memory.memory) {
		return 0x0, errors.New("out of bound read")
	}
	memory.logger.Debug("read byte from memory", "address", address, "value", memory.memory[int(address)])
	return memory.memory[int(address)], nil
}

// ReadDouble reads a uint16 at the given address and returns it.
func (memory *Memory) ReadDouble(address uint16) (uint16, error) {
	if int(address)+2 > len(memory.memory) {
		return 0x0, errors.New("out of bound read")
	}
	lo, err := memory.ReadByte(address)
	if err != nil {
		return 0x0, err
	}
	hi, err := memory.ReadByte(address + 1)
	if err != nil {
		return 0x0, err
	}
	return uint16(hi)<<8 | uint16(lo), nil
}

// Read reads into the buffer the memory at the provided address until buffer is full.
func (memory *Memory) Read(address uint16, buffer *[]byte) error {
	if int(address)+len(*buffer) > len(memory.memory) {
		return errors.New("out of bound read")
	}
	copy(*buffer, memory.memory[int(address):])
	return nil
}

// ReaderAt creates and returns a bufio.Reader at a given memory address.
func (memory *Memory) ReaderAt(address uint16) (*bufio.Reader, error) {
	if int(address) > len(memory.memory) {
		return nil, errors.New("out of bound read")
	}
	return bufio.NewReader(bytes.NewReader(memory.memory[int(address):])), nil
}

// Dump returns a hexdump of the memory
func (memory *Memory) Dump() string {
	return hex.Dump(memory.memory)
}
