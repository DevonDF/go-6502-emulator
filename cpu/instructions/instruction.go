package instructions

import (
	"emulator/cpu"
	"emulator/memory"
)

type InstructionHandler interface {
	Execute(*cpu.CPU, *memory.Memory, *DecodedInstruction) // execute the instruction on the given CPU & Memory.
}

// type Instruction struct {
// 	opcode  byte               // the opcode for the given instruction.
// 	size    byte               // the number of bytes this instruction takes.
// 	cycles  byte               // the number of cpu cycles this operation takes.
// 	handler InstructionHandler // the handler for this instruction.
// }
