package instructions

// AddressingMode is a byte backed enum for addressing mode of an instruction.
type AddressingMode byte

const (
	AddrAccumulator AddressingMode = iota
	AddrAbsolute
	AddrAbsoluteX
	AddrAbsoluteY
	AddrImmediate
	AddrImplied
	AddrIndirect
	AddrIndirectX
	AddrIndirectY
	AddrRelative
	AddrZeropage
	AddrZeropageX
	AddrZeropageY
)

// Instruction defines an instruction within the 6502 instruction set.
type Instruction struct {
	AssemblyString string         // the assembly string for this instruction, e.g. AND, ADD, etc.
	Opcode         byte           // the opcode for the given instruction.
	Size           byte           // the number of bytes this instruction takes.
	AddressingMode AddressingMode // the addressing mode for this instruction.
	Cycles         byte           // the number of cpu cycles this operation takes.
}

var instructionSet = [256]Instruction{
	// ADC - Add Memory to Accumulator with Carry
	0x69: { // ADC #oper
		AssemblyString: "ADC",
		Opcode:         0x69,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
	},
	0x65: { // ADC oper
		AssemblyString: "ADC",
		Opcode:         0x65,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
	},
	0x75: { // ADC oper,X
		AssemblyString: "ADC",
		Opcode:         0x75,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
	},
	0x6D: { // ADC oper
		AssemblyString: "ADC",
		Opcode:         0x6D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
	},
	0x7D: { // ADC oper,X
		AssemblyString: "ADC",
		Opcode:         0x7D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
	},
	0x79: { // ADC oper,Y
		AssemblyString: "ADC",
		Opcode:         0x79,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
	},
	0x61: { // ADC (oper,X)
		AssemblyString: "ADC",
		Opcode:         0x61,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
	},
	0x71: { // ADC (oper),Y
		AssemblyString: "ADC",
		Opcode:         0x71,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
	},

	// AND - AND Memory with Accumulator
	0x29: { // AND #oper
		AssemblyString: "AND",
		Opcode:         0x29,
		Size:           2,
		Cycles:         2,
	},
	0x25: { // AND oper
		AssemblyString: "AND",
		Opcode:         0x25,
		Size:           2,
		Cycles:         3,
	},
	0x35: { // AND oper,X
		AssemblyString: "AND",
		Opcode:         0x35,
		Size:           2,
		Cycles:         4,
	},
	0x2D: { // AND oper
		AssemblyString: "AND",
		Opcode:         0x2D,
		Size:           3,
		Cycles:         4,
	},
	0x3D: { // AND oper,X
		AssemblyString: "AND",
		Opcode:         0x7D,
		Size:           3,
		Cycles:         4,
	},
	0x39: { // AND oper,Y
		AssemblyString: "AND",
		Opcode:         0x39,
		Size:           3,
		Cycles:         4,
	},
	0x21: { // AND (oper,X)
		AssemblyString: "AND",
		Opcode:         0x21,
		Size:           2,
		Cycles:         6,
	},
	0x31: { // AND (oper),Y
		AssemblyString: "AND",
		Opcode:         0x31,
		Size:           2,
		Cycles:         5,
	},

	// ASL - Shift Left One Bit (Memory or Accumulator)
	0x0A: { // ASL A
		Opcode: 0x0A,
		Size:   1,
		Cycles: 2,
	},
	0x06: { // ASL oper
		Opcode: 0x06,
		Size:   2,
		Cycles: 5,
	},
	0x16: { // ASL oper,X
		Opcode: 0x16,
		Size:   2,
		Cycles: 6,
	},
	0x0E: { // ASL oper
		Opcode: 0x0E,
		Size:   3,
		Cycles: 6,
	},
	0x1E: { // ASL oper,X
		Opcode: 0x1E,
		Size:   3,
		Cycles: 7,
	},

	// BCC - Branch on Carry Clear
	0x90: {
		Opcode: 0x90,
		Size:   2,
		Cycles: 2,
	},

	// BCS - Branch on Carry Set
	0xB0: {
		Opcode: 0xB0,
		Size:   2,
		Cycles: 2,
	},

	// BEQ - Branch on Result Zero
	0xF0: {
		Opcode: 0xF0,
		Size:   2,
		Cycles: 2,
	},

	// BIT - Test Bits in Memory with Accumulator
	0x24: {
		Opcode: 0x24,
		Size:   2,
		Cycles: 3,
	},
	0x2C: {
		Opcode: 0x2C,
		Size:   2,
		Cycles: 3,
	},

	// BMI - Branch on Result Minus
	0x30: {
		Opcode: 0x30,
		Size:   2,
		Cycles: 2,
	},

	// BNE - Branch on Result not Zero
	0xD0: {
		Opcode: 0xD0,
		Size:   2,
		Cycles: 2,
	},

	// BPL - Branch on Result Plus
	0x10: {
		Opcode: 0x10,
		Size:   2,
		Cycles: 2,
	},

	// BRK - Break & Interrupt
	0x00: { // BRK
		Opcode: 0x00,
		Size:   1,
		Cycles: 7,
	},

	// BVC - Branch on Overflow Clear
	0x50: {
		Opcode: 0x50,
		Size:   2,
		Cycles: 2,
	},

	// BVS - Branch on Overflow Set
	0x70: {
		Opcode: 0x70,
		Size:   2,
		Cycles: 2,
	},

	// LDA - Load Accumulator with Memory
	0xA9: {
		Opcode: 0xA9,
		Size:   2,
		Cycles: 2,
	},
	0xA5: {
		Opcode: 0xA5,
		Size:   2,
		Cycles: 3,
	},
	0xB5: {
		Opcode: 0xB5,
		Size:   2,
		Cycles: 4,
	},
	0xAD: {
		Opcode: 0xAD,
		Size:   3,
		Cycles: 4,
	},
	0xBD: {
		Opcode: 0xBD,
		Size:   3,
		Cycles: 4,
	},
	0xB9: {
		Opcode: 0xB9,
		Size:   3,
		Cycles: 4,
	},
	0xA1: {
		Opcode: 0xA1,
		Size:   2,
		Cycles: 6,
	},
	0xA2: {
		Opcode: 0xA2,
		Size:   2,
		Cycles: 5,
	},

	// STA - Store Accumulator in Memory
	0x85: {
		Opcode: 0x85,
		Size:   2,
		Cycles: 3,
	},
	0x95: {
		Opcode: 0xA5,
		Size:   2,
		Cycles: 4,
	},
	0x8D: {
		Opcode: 0xB5,
		Size:   3,
		Cycles: 4,
	},
	0x9D: {
		Opcode: 0xAD,
		Size:   3,
		Cycles: 5,
	},
	0x99: {
		Opcode: 0xBD,
		Size:   3,
		Cycles: 5,
	},
	0x81: {
		Opcode: 0xB9,
		Size:   2,
		Cycles: 6,
	},
	0x91: {
		Opcode: 0xA1,
		Size:   2,
		Cycles: 6,
	},
}

// InstructionFromOpcode returns a pointer to an Instruction from a provided opcode.
func InstructionFromOpcode(opcode byte) *Instruction {
	return &instructionSet[opcode]
}

// InstructionFromAssembly returns an Instruction pointer from an assembly string and addressing mode.
func InstructionFromAssembly(assemblyString string, addressingMode AddressingMode) *Instruction {
	for _, inst := range instructionSet {
		if inst.AssemblyString == assemblyString && inst.AddressingMode == addressingMode {
			return &inst
		}
	}
	return nil
}
