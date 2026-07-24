package core

import (
	"github.com/adityamkk/cc-isa-emulator/isa"
)

func (core *Core) execute(instr *isa.Instruction) (uint16, bool, error) {
	if instr.Raw == 0 {
		return core.registers.Pc + 2, true, nil
	}
	return core.registers.Pc + 2, false, nil
}
