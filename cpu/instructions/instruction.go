package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type InstructionHandler interface {
	Execute(*cpu.CPU, *memory.Memory, *DecodedInstruction) error // execute the instruction on the given CPU & Memory.
}
