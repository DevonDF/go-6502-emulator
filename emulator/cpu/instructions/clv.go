package instructions

import (
	"github.com/DevonDF/go-6502-emulator/emulator/cpu"
	"github.com/DevonDF/go-6502-emulator/emulator/memory"
)

type CLVHandler struct {
}

var clvHandler = &CLVHandler{}

func CLV() *CLVHandler {
	return clvHandler
}

func (handler *CLVHandler) Execute(cpu *cpu.CPU, memory *memory.Memory, instruction *DecodedInstruction) error {
	// 0 -> V
	cpu.SetOverflowFlag(false)
	return nil
}
