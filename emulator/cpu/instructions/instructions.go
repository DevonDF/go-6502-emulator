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

	// CLC - Clear Carry Flag
	0x18: {
		AssemblyString: "CLC",
		Opcode:         0x18,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        CLC(),
	},

	// CLD - Clear Decimal Mode
	0xD8: {
		AssemblyString: "CLD",
		Opcode:         0xD8,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        CLD(),
	},

	// CLI - Clear Interrupt Disable Bit
	0x58: {
		AssemblyString: "CLI",
		Opcode:         0x58,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        CLI(),
	},

	// CLV - Clear Overflow Flag
	0xB8: {
		AssemblyString: "CLV",
		Opcode:         0xB8,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        CLV(),
	},

	// CMP - Compare Memory with Accumulator
	0xC9: {
		AssemblyString: "CMP",
		Opcode:         0xC9,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        CMP(),
	},
	0xC5: {
		AssemblyString: "CMP",
		Opcode:         0xC5,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        CMP(),
	},
	0xD5: {
		AssemblyString: "CMP",
		Opcode:         0xD5,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        CMP(),
	},
	0xCD: {
		AssemblyString: "CMP",
		Opcode:         0xCD,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        CMP(),
	},
	0xDD: {
		AssemblyString: "CMP",
		Opcode:         0xDD,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        CMP(),
	},
	0xD9: {
		AssemblyString: "CMP",
		Opcode:         0xD9,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        CMP(),
	},
	0xC1: {
		AssemblyString: "CMP",
		Opcode:         0xC1,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        CMP(),
	},
	0xD1: {
		AssemblyString: "CMP",
		Opcode:         0xD1,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        CMP(),
	},

	// CPX - Compare Memory with Register X
	0xE0: {
		AssemblyString: "CPX",
		Opcode:         0xE0,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        CPX(),
	},
	0xE4: {
		AssemblyString: "CPX",
		Opcode:         0xE4,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        CPX(),
	},
	0xEC: {
		AssemblyString: "CPX",
		Opcode:         0xEC,
		Size:           2,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        CPX(),
	},

	// CPY - Compare Memory with Register X
	0xC0: {
		AssemblyString: "CPY",
		Opcode:         0xC0,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        CPY(),
	},
	0xC4: {
		AssemblyString: "CPY",
		Opcode:         0xC4,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        CPY(),
	},
	0xCC: {
		AssemblyString: "CPY",
		Opcode:         0xCC,
		Size:           2,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        CPY(),
	},

	// DEC - Decrement Memory by One
	0xC6: {
		AssemblyString: "DEC",
		Opcode:         0xC6,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        DEC(),
	},
	0xD6: {
		AssemblyString: "DEC",
		Opcode:         0xD6,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        DEC(),
	},
	0xCE: {
		AssemblyString: "DEC",
		Opcode:         0xCE,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        DEC(),
	},
	0xDE: {
		AssemblyString: "DEC",
		Opcode:         0xDE,
		Size:           2,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        DEC(),
	},

	// DEX - Decrement Register X by One
	0xCA: {
		AssemblyString: "DEX",
		Opcode:         0xCA,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        DEX(),
	},

	// DEY - Decrement Register Y by One
	0x88: {
		AssemblyString: "DEY",
		Opcode:         0x88,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        DEY(),
	},

	// EOR - Exclusive-OR Memory with Accumulator
	0x49: {
		AssemblyString: "EOR",
		Opcode:         0x49,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        EOR(),
	},
	0x45: {
		AssemblyString: "EOR",
		Opcode:         0x45,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        EOR(),
	},
	0x55: {
		AssemblyString: "EOR",
		Opcode:         0x55,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        EOR(),
	},
	0x4D: {
		AssemblyString: "EOR",
		Opcode:         0x4D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        EOR(),
	},
	0x5D: {
		AssemblyString: "EOR",
		Opcode:         0x5D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        EOR(),
	},
	0x59: {
		AssemblyString: "EOR",
		Opcode:         0x59,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        EOR(),
	},
	0x41: {
		AssemblyString: "EOR",
		Opcode:         0x41,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        EOR(),
	},
	0x51: {
		AssemblyString: "EOR",
		Opcode:         0x51,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        EOR(),
	},

	// INC - Increment Memory by One
	0xE6: {
		AssemblyString: "INC",
		Opcode:         0xE6,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        INC(),
	},
	0xF6: {
		AssemblyString: "INC",
		Opcode:         0xF6,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        INC(),
	},
	0xEE: {
		AssemblyString: "INC",
		Opcode:         0xEE,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        INC(),
	},
	0xFE: {
		AssemblyString: "INC",
		Opcode:         0xFE,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        INC(),
	},

	// INX - Increment Index X by One
	0xE8: {
		AssemblyString: "INX",
		Opcode:         0xE8,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        INX(),
	},

	// INY - Increment Index Y by One
	0xC8: {
		AssemblyString: "INY",
		Opcode:         0xC8,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        INY(),
	},

	// JMP - Jump to New Location
	0x4C: {
		AssemblyString: "JMP",
		Opcode:         0x4C,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         3,
		Handler:        JMP(),
	},
	0x6C: {
		AssemblyString: "JMP",
		Opcode:         0x6C,
		Size:           3,
		AddressingMode: AddrIndirect,
		Cycles:         5,
		Handler:        JMP(),
	},

	// JSR - Jump to New Location Saving Return Address
	0x20: {
		AssemblyString: "JSR",
		Opcode:         0x20,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        JSR(),
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
	0xB1: {
		AssemblyString: "LDA",
		Opcode:         0xB1,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        LDA(),
	},

	// LDX - Load Index X with Memory
	0xA2: {
		AssemblyString: "LDX",
		Opcode:         0xA2,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        LDX(),
	},
	0xA6: {
		AssemblyString: "LDX",
		Opcode:         0xA6,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        LDX(),
	},
	0xB6: {
		AssemblyString: "LDX",
		Opcode:         0xB6,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        LDX(),
	},
	0xAE: {
		AssemblyString: "LDX",
		Opcode:         0xAE,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        LDX(),
	},
	0xBE: {
		AssemblyString: "LDX",
		Opcode:         0xBE,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        LDX(),
	},

	// LDY - Load Index Y with Memory
	0xA0: {
		AssemblyString: "LDY",
		Opcode:         0xA0,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        LDY(),
	},
	0xA4: {
		AssemblyString: "LDY",
		Opcode:         0xA4,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        LDY(),
	},
	0xB4: {
		AssemblyString: "LDY",
		Opcode:         0xB4,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        LDY(),
	},
	0xAC: {
		AssemblyString: "LDY",
		Opcode:         0xAC,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        LDY(),
	},
	0xBC: {
		AssemblyString: "LDY",
		Opcode:         0xBC,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        LDY(),
	},

	// LSR - Shift One Bit Right
	0x4A: {
		AssemblyString: "LSR",
		Opcode:         0x4A,
		Size:           1,
		AddressingMode: AddrAccumulator,
		Cycles:         2,
		Handler:        LSR(),
	},
	0x46: {
		AssemblyString: "LSR",
		Opcode:         0x46,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        LSR(),
	},
	0x56: {
		AssemblyString: "LSR",
		Opcode:         0x56,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        LSR(),
	},
	0x4E: {
		AssemblyString: "LSR",
		Opcode:         0x4E,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        LSR(),
	},
	0x5E: {
		AssemblyString: "LSR",
		Opcode:         0x5E,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        LSR(),
	},

	// NOP - No Operation
	0xEA: {
		AssemblyString: "NOP",
		Opcode:         0xEA,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        NOP(),
	},

	// ORA - OR Memory with Accumulator
	0x09: {
		AssemblyString: "ORA",
		Opcode:         0x09,
		Size:           2,
		AddressingMode: AddrImmediate,
		Cycles:         2,
		Handler:        ORA(),
	},
	0x05: {
		AssemblyString: "ORA",
		Opcode:         0x05,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        ORA(),
	},
	0x15: {
		AssemblyString: "ORA",
		Opcode:         0x15,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         4,
		Handler:        ORA(),
	},
	0x0D: {
		AssemblyString: "ORA",
		Opcode:         0x0D,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        ORA(),
	},
	0x1D: {
		AssemblyString: "ORA",
		Opcode:         0x1D,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         4,
		Handler:        ORA(),
	},
	0x19: {
		AssemblyString: "ORA",
		Opcode:         0x19,
		Size:           3,
		AddressingMode: AddrAbsoluteY,
		Cycles:         4,
		Handler:        ORA(),
	},
	0x01: {
		AssemblyString: "ORA",
		Opcode:         0x01,
		Size:           2,
		AddressingMode: AddrIndirectX,
		Cycles:         6,
		Handler:        ORA(),
	},
	0x11: {
		AssemblyString: "ORA",
		Opcode:         0x11,
		Size:           2,
		AddressingMode: AddrIndirectY,
		Cycles:         5,
		Handler:        ORA(),
	},

	// PHA - Push Accumulator on Stack
	0x48: {
		AssemblyString: "PHA",
		Opcode:         0x48,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         3,
		Handler:        PHA(),
	},

	// PHP - Push Processor Status on Stack
	0x08: {
		AssemblyString: "PHP",
		Opcode:         0x08,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         3,
		Handler:        PHP(),
	},

	// PLA - Pull Accumulator from Stack
	0x68: {
		AssemblyString: "PLA",
		Opcode:         0x68,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         4,
		Handler:        PLA(),
	},

	// PLP - Pull Processor Status from Stack
	0x28: {
		AssemblyString: "PLP",
		Opcode:         0x28,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         4,
		Handler:        PLP(),
	},

	// ROL - Rotate One Bit Left
	0x2A: {
		AssemblyString: "ROL",
		Opcode:         0x2A,
		Size:           1,
		AddressingMode: AddrAccumulator,
		Cycles:         2,
		Handler:        ROL(),
	},
	0x26: {
		AssemblyString: "ROL",
		Opcode:         0x26,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        ROL(),
	},
	0x36: {
		AssemblyString: "ROL",
		Opcode:         0x36,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        ROL(),
	},
	0x2E: {
		AssemblyString: "ROL",
		Opcode:         0x2E,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        ROL(),
	},
	0x3E: {
		AssemblyString: "ROL",
		Opcode:         0x3E,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        ROL(),
	},

	// ROR - Rotate One Bit Right
	0x6A: {
		AssemblyString: "ROR",
		Opcode:         0x6A,
		Size:           1,
		AddressingMode: AddrAccumulator,
		Cycles:         2,
		Handler:        ROR(),
	},
	0x66: {
		AssemblyString: "ROR",
		Opcode:         0x66,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         5,
		Handler:        ROR(),
	},
	0x76: {
		AssemblyString: "ROR",
		Opcode:         0x76,
		Size:           2,
		AddressingMode: AddrZeropageX,
		Cycles:         6,
		Handler:        ROR(),
	},
	0x6E: {
		AssemblyString: "ROR",
		Opcode:         0x6E,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         6,
		Handler:        ROR(),
	},
	0x7E: {
		AssemblyString: "ROR",
		Opcode:         0x7E,
		Size:           3,
		AddressingMode: AddrAbsoluteX,
		Cycles:         7,
		Handler:        ROR(),
	},

	// RTI - Return from Interrupt
	0x40: {
		AssemblyString: "RTI",
		Opcode:         0x40,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         6,
		Handler:        RTI(),
	},

	// RTS - Return from Subroutine
	0x60: {
		AssemblyString: "RTS",
		Opcode:         0x60,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         6,
		Handler:        RTS(),
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

	// STX - Store Index X in Memory
	0x86: {
		AssemblyString: "STX",
		Opcode:         0x86,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        STX(),
	},
	0x96: {
		AssemblyString: "STX",
		Opcode:         0x96,
		Size:           2,
		AddressingMode: AddrZeropageY,
		Cycles:         4,
		Handler:        STX(),
	},
	0x8E: {
		AssemblyString: "STX",
		Opcode:         0x8E,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        STX(),
	},

	// STY - Store Index Y in Memory
	0x84: {
		AssemblyString: "STY",
		Opcode:         0x84,
		Size:           2,
		AddressingMode: AddrZeropage,
		Cycles:         3,
		Handler:        STY(),
	},
	0x94: {
		AssemblyString: "STY",
		Opcode:         0x94,
		Size:           2,
		AddressingMode: AddrZeropageY,
		Cycles:         4,
		Handler:        STY(),
	},
	0x8C: {
		AssemblyString: "STY",
		Opcode:         0x8C,
		Size:           3,
		AddressingMode: AddrAbsolute,
		Cycles:         4,
		Handler:        STY(),
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

	// TAY - Transfer Accumulator to Index Y
	0xA8: {
		AssemblyString: "TAY",
		Opcode:         0xA8,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TAY(),
	},

	// TSX - Transfer Stack Pointer to Index X
	0xBA: {
		AssemblyString: "TSX",
		Opcode:         0xBA,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TSX(),
	},

	// TXA - Transfer Index X to Accumulator
	0x8A: {
		AssemblyString: "TXA",
		Opcode:         0x8A,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TXA(),
	},

	// TXS - Transfer Index X to Stack Register
	0x9A: {
		AssemblyString: "TXS",
		Opcode:         0x9A,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TXS(),
	},

	// TYA - Transfer Index Y to Accumulator
	0x98: {
		AssemblyString: "TYA",
		Opcode:         0x98,
		Size:           1,
		AddressingMode: AddrImplied,
		Cycles:         2,
		Handler:        TYA(),
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

// LargestInstructionSizeForOpcode returns the largest possible instruction size for a given opcode string.
func LargestInstructionSizeForOpcode(assemblyString string) int {
	size := 0
	for _, inst := range instructionSet {
		if inst.AssemblyString == assemblyString {
			size = max(size, int(inst.Size))
		}
	}
	return size
}
