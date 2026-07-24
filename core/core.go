package core

import (
	"github.com/adityamkk/cc-isa-emulator/isa"
	"github.com/adityamkk/cc-isa-emulator/mem"
)

const NUM_REGISTERS = 16

type Registers struct {
	R  [NUM_REGISTERS]uint16
	Pc uint16
}

func NewRegisters() *Registers {
	return &Registers{R: [16]uint16{}, Pc: uint16(0)}
}

type Core struct {
	registers *Registers
	ram       *mem.RandomAccessMemory
	Stop      chan error
}

func NewCore(ram *mem.RandomAccessMemory) *Core {
	return &Core{registers: NewRegisters(), ram: ram, Stop: make(chan error)}
}

func (core *Core) fetch(address uint16) uint16 {
	return core.ram.Read1Word(address)
}

func (core *Core) Run() {
	var stop bool = false
	for !stop {
		instr := core.fetch(core.registers.Pc)
		pcNew, done, err := core.execute(isa.NewInstruction(instr))
		if err != nil {
			core.Stop <- err
			return
		}
		core.registers.Pc = pcNew
		stop = done
	}
	core.Stop <- nil
	return
}
