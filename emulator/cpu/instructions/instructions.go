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
	AddrRelative
	AddrIndirect
	AddrIndirectX
	AddrIndirectY
	AddrZeropage
	AddrZeropageX
	AddrZeropageY
)

// Instruction defines an instruction within the 6502 instruction set.
type Instruction struct {
	AssemblyString string             // the assembly string for this instruction, e.g. AND, ADD, etc.
	Opcode         byte               // the opcode for the given instruction.
	Size           byte               // the number of bytes this instruction takes.
	AddressingMode AddressingMode     // the addressing mode for this instruction.
	Cycles         byte               // the number of cpu cycles this operation takes.
	Handler        InstructionHandler // the handler to execute this instruction
}

var instructionSet = [256]Instruction{
	// ADC - Add Memory to Accumulator with Carry
	0x69: { // ADC #oper
		AssemblyString: "ADC",
		Opcode:         0x69,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        ADC(),
	},
	0x65: { // ADC oper
		AssemblyString: "ADC",
		Opcode:         0x65,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        ADC(),
	},
	0x75: { // ADC oper,X
		AssemblyString: "ADC",
		Opcode:         0x75,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        ADC(),
	},
	0x6D: { // ADC oper
		AssemblyString: "ADC",
		Opcode:         0x6D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        ADC(),
	},
	0x7D: { // ADC oper,X
		AssemblyString: "ADC",
		Opcode:         0x7D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        ADC(),
	},
	0x79: { // ADC oper,Y
		AssemblyString: "ADC",
		Opcode:         0x79,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        ADC(),
	},
	0x61: { // ADC (oper,X)
		AssemblyString: "ADC",
		Opcode:         0x61,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        ADC(),
	},
	0x71: { // ADC (oper),Y
		AssemblyString: "ADC",
		Opcode:         0x71,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        ADC(),
	},

	// AND - AND Memory with Accumulator
	0x29: { // AND #oper
		AssemblyString: "AND",
		Opcode:         0x29,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        AND(),
	},
	0x25: { // AND oper
		AssemblyString: "AND",
		Opcode:         0x25,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        AND(),
	},
	0x35: { // AND oper,X
		AssemblyString: "AND",
		Opcode:         0x35,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        AND(),
	},
	0x2D: { // AND oper
		AssemblyString: "AND",
		Opcode:         0x2D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        AND(),
	},
	0x3D: { // AND oper,X
		AssemblyString: "AND",
		Opcode:         0x7D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        AND(),
	},
	0x39: { // AND oper,Y
		AssemblyString: "AND",
		Opcode:         0x39,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        AND(),
	},
	0x21: { // AND (oper,X)
		AssemblyString: "AND",
		Opcode:         0x21,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        AND(),
	},
	0x31: { // AND (oper),Y
		AssemblyString: "AND",
		Opcode:         0x31,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        AND(),
	},

	// ASL - Shift Left One Bit (Memory or Accumulator)
	0x0A: { // ASL A
		AssemblyString: "ASL",
		Opcode:         0x0A,
		Size:           1,
		AddressingMode: AddrAccumulator,
		Cycles:         2,
		Handler:        ASL(),
	},
	0x06: { // ASL oper
		AssemblyString: "ASL",
		Opcode:         0x06,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        ASL(),
	},
	0x16: { // ASL oper,X
		AssemblyString: "ASL",
		Opcode:         0x16,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        ASL(),
	},
	0x0E: { // ASL oper
		AssemblyString: "ASL",
		Opcode:         0x0E,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        ASL(),
	},
	0x1E: { // ASL oper,X
		AssemblyString: "ASL",
		Opcode:         0x1E,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        ASL(),
	},

	// BCC - Branch on Carry Clear
	0x90: {
		AssemblyString: "BCC",
		Opcode:         0x90,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BCC(),
	},

	// BCS - Branch on Carry Set
	0xB0: {
		AssemblyString: "BCS",
		Opcode:         0xB0,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BCS(),
	},

	// BEQ - Branch on Result Zero
	0xF0: {
		AssemblyString: "BEQ",
		Opcode:         0xF0,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BEQ(),
	},

	// BIT - Test Bits in Memory with Accumulator
	0x24: {
		AssemblyString: "BIT",
		Opcode:         0x24,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        Unimpl(),
	},
	0x2C: {
		AssemblyString: "BIT",
		Opcode:         0x2C,
		Size:           2,
		AddressingMode: AddrAbsolute,
		Cycles:         3,
		Handler:        Unimpl(),
	},

	// BMI - Branch on Result Minus
	0x30: {
		AssemblyString: "BMI",
		Opcode:         0x30,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BMI(),
	},

	// BNE - Branch on Result not Zero
	0xD0: {
		AssemblyString: "BNE",
		Opcode:         0xD0,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BNE(),
	},

	// BPL - Branch on Result Plus
	0x10: {
		AssemblyString: "BPL",
		Opcode:         0x10,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BPL(),
	},

	// BRK - Break & Interrupt
	0x00: { // BRK
		AssemblyString: "BRK",
		Opcode:         0x00,
		AddressingMode: AddrImplied,
		Size:           1,
		Cycles:         7,
		Handler:        BRK(),
	},

	// BVC - Branch on Overflow Clear
	0x50: {
		AssemblyString: "BVC",
		Opcode:         0x50,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BVC(),
	},

	// BVS - Branch on Overflow Set
	0x70: {
		AssemblyString: "BVS",
		Opcode:         0x70,
		Size:           2,
		AddressingMode: AddrRelative,
		Cycles:         2,
		Handler:        BVS(),
	},

	// LDA - Load Accumulator with Memory
	0xA9: {
		AssemblyString: "LDA",
		Opcode:         0xA9,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        LDA(),
	},
	0xA5: {
		AssemblyString: "LDA",
		Opcode:         0xA5,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        LDA(),
	},
	0xB5: {
		AssemblyString: "LDA",
		Opcode:         0xB5,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        LDA(),
	},
	0xAD: {
		AssemblyString: "LDA",
		Opcode:         0xAD,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        LDA(),
	},
	0xBD: {
		AssemblyString: "LDA",
		Opcode:         0xBD,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        LDA(),
	},
	0xB9: {
		AssemblyString: "LDA",
		Opcode:         0xB9,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        LDA(),
	},
	0xA1: {
		AssemblyString: "LDA",
		Opcode:         0xA1,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        LDA(),
	},
	0xA2: {
		AssemblyString: "LDA",
		Opcode:         0xA2,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        LDA(),
	},

	// STA - Store Accumulator in Memory
	0x85: {
		AssemblyString: "STA",
		Opcode:         0x85,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        STA(),
	},
	0x95: {
		AssemblyString: "STA",
		Opcode:         0x95,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        STA(),
	},
	0x8D: {
		AssemblyString: "STA",
		Opcode:         0x8D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        STA(),
	},
	0x9D: {
		AssemblyString: "STA",
		Opcode:         0x9D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         5,
		Handler:        STA(),
	},
	0x99: {
		AssemblyString: "STA",
		Opcode:         0x99,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         5,
		Handler:        STA(),
	},
	0x81: {
		AssemblyString: "STA",
		Opcode:         0x81,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        STA(),
	},
	0x91: {
		AssemblyString: "STA",
		Opcode:         0x91,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         6,
		Handler:        STA(),
	},

	// TAX - Transfer Accumulator to Index X
	0xAA: {
		AssemblyString: "TAX",
		Opcode:         0xAA,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TAX(),
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
