package core

import (
	"github.com/adityamkk/cc-isa-emulator/isa"
)

func (core *Core) execute(instr *isa.Instruction) (uint16, bool, error) {
	return uint16(0), true, nil
}
