package addressing

import (
	"emulator/cpu"
	"emulator/memory"
)

func GetZeropageAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[0]), nil
}

func GetZeropageXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[0]) + uint16(cpu.RegisterX), nil
}

func GetAbsoluteAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[1])<<8 | uint16(operands[0]), nil
}

func GetAbsoluteXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return (uint16(operands[1])<<8 | uint16(operands[0])) + uint16(cpu.RegisterX), nil
}

func GetAbsoluteYAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0])
	addr2, err := memory.Read16(addr)
	return addr2 + uint16(cpu.RegisterY), err
}

func GetIndirectXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0]) + uint16(cpu.RegisterX)
	addr2, err := memory.Read16(addr)
	return addr2, err
}

func GetIndirectYAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0]) + uint16(cpu.RegisterX)
	addr2, err := memory.Read16(addr)
	return addr2, err
}

// Read a byte from memory using zeropage addressing.
func ReadZeropage(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetZeropageAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using zeropage,X addressing.
func ReadZeropageX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetZeropageXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using absolute addressing.
func ReadAbsolute(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using absolute,X addressing.
func ReadAbsoluteX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using absolute,Y addressing.
func ReadAbsoluteY(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteYAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using (indirect,X) addressing.
func ReadIndirectX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetIndirectXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// Read a byte from memory using (indirect,Y) addressing.
func ReadIndirectY(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteYAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}
