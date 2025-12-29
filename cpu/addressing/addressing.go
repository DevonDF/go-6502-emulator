package addressing

import (
	"emulator/cpu"
	"emulator/memory"
)

// GetZeropageAddress returns the address in memory given by the provided operands using zeropage addressing.
func GetZeropageAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[0]), nil
}

// GetZeropageXAddress returns the address in memory given by the provided operands using zeropage,X addressing.
func GetZeropageXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[0]) + uint16(cpu.RegisterX), nil
}

// GetAbsoluteAddress returns the address in memory given by the provided operands using absolute addressing.
func GetAbsoluteAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return uint16(operands[1])<<8 | uint16(operands[0]), nil
}

// GetAbsoluteXAddress returns the address in memory given by the provided operands using absolute,X addressing.
func GetAbsoluteXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	return (uint16(operands[1])<<8 | uint16(operands[0])) + uint16(cpu.RegisterX), nil
}

// GetAbsoluteYAddress returns the address in memory given by the provided operands using absolute,Y addressing.
func GetAbsoluteYAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0])
	addr2, err := memory.Read16(addr)
	return addr2 + uint16(cpu.RegisterY), err
}

// GetIndirectXAddress returns the address in memory given by the provided operands using indirect,X addressing.
func GetIndirectXAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0]) + uint16(cpu.RegisterX)
	addr2, err := memory.Read16(addr)
	return addr2, err
}

// GetIndirectYAddress returns the address in memory given by the provided operands using indirect,Y addressing.
func GetIndirectYAddress(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (uint16, error) {
	addr := uint16(operands[0]) + uint16(cpu.RegisterX)
	addr2, err := memory.Read16(addr)
	return addr2, err
}

// ReadZeropage reads a byte from memory using zeropage addressing.
func ReadZeropage(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetZeropageAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadZeropageX reads a byte from memory using zeropage,X addressing.
func ReadZeropageX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetZeropageXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadAbsolute reads a byte from memory using absolute addressing.
func ReadAbsolute(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadAbsoluteX reads a byte from memory using absolute,X addressing.
func ReadAbsoluteX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadAbsoluteY reads a byte from memory using absolute,Y addressing.
func ReadAbsoluteY(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteYAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadIndirectX reads a byte from memory using (indirect,X) addressing.
func ReadIndirectX(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetIndirectXAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}

// ReadIndirectY reads a byte from memory using (indirect,Y) addressing.
func ReadIndirectY(operands []byte, cpu *cpu.CPU, memory *memory.Memory) (byte, error) {
	addr, err := GetAbsoluteYAddress(operands, cpu, memory)
	if err != nil {
		return 0x0, err
	}
	return memory.ReadByte(addr)
}
