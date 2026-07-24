package isa

type Instruction struct {
	Raw uint16
}

func NewInstruction(instr uint16) *Instruction {
	return &Instruction{Raw: instr}
}
