package cpu

import (
	"log/slog"

	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type Stack struct {
	cpu    *CPU
	logger *slog.Logger
}

func NewStack(cpu *CPU, logger *slog.Logger) *Stack {
	return &Stack{
		cpu:    cpu,
		logger: logger,
	}
}

func (stack *Stack) getStackPointerAddress() uint16 {
	return (memory.StackStartAddress + uint16(stack.cpu.RegisterSP))
}

// PushByte pushes a byte onto the stack.
func (stack *Stack) PushByte(value int8, memory_ *memory.Memory) error {
	stack.logger.Debug("stack PushByte", "value", value, "SP", stack.cpu.RegisterSP, "NewSP", stack.cpu.RegisterSP+1)

	err := memory_.Write(stack.getStackPointerAddress(), []byte{byte(value)})
	if err != nil {
		return err
	}

	stack.cpu.RegisterSP = stack.cpu.RegisterSP + 1
	return nil
}

// PushDouble pushes a double onto the stack.
func (stack *Stack) PushDouble(value int16, memory_ *memory.Memory) error {
	stack.logger.Debug("stack PushByte", "value", value, "SP", stack.cpu.RegisterSP, "NewSP", stack.cpu.RegisterSP+2)

	err := memory_.Write(stack.getStackPointerAddress(), []byte{byte(value & 0xFF), byte((value >> 8) & 0xFF)})
	if err != nil {
		return err
	}

	stack.cpu.RegisterSP = stack.cpu.RegisterSP + 2
	return nil
}

// PopByte pops a byte from the stack.
func (stack *Stack) PopByte(memory_ *memory.Memory) (int8, error) {
	stack.cpu.RegisterSP = stack.cpu.RegisterSP - 1
	readByte, err := memory_.ReadByte(stack.getStackPointerAddress())
	if err != nil {
		return 0x0, err
	}

	stack.logger.Debug("stack PushByte", "value", readByte, "SP", stack.cpu.RegisterSP+1, "NewSP", stack.cpu.RegisterSP)
	return int8(readByte), nil
}

// PopDouble pops a double from the stack.
func (stack *Stack) PopDouble(memory_ *memory.Memory) (int16, error) {
	stack.cpu.RegisterSP = stack.cpu.RegisterSP - 2
	readDouble, err := memory_.ReadDouble(stack.getStackPointerAddress())
	if err != nil {
		return 0x0, err
	}

	stack.logger.Debug("stack PushByte", "value", readDouble, "SP", stack.cpu.RegisterSP+2, "NewSP", stack.cpu.RegisterSP)
	return int16(readDouble), nil
}
