package assembler

type AssemblyInstructionLabel struct {
	labelName string // the name of the label.
	RawAssemblyInstruction
}

func (label *AssemblyInstructionLabel) ToByteCode(assembler *Assembler) ([]byte, error) {
	return []byte{}, nil
}
