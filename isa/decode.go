package isa

type Instruction struct {
	raw uint16
}

func NewInstruction(instr uint16) *Instruction {
	return &Instruction{raw: instr}
}
