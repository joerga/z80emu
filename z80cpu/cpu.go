package z80cpu

import (
	"bytes"
	"strconv"
	"z80emu/Memory"
)

// R16 represents a 16-bit register

// R8 represents a 8-bit register
type R8 byte
type z80Instruction byte

const (
	Carry = iota
	Null
	Overflow
	HalfCarry
	Zero
	Sign
)

type (
	execInstruction func(val *Z80)
)

// Z80 contains all internal registers and such for the Z80 processor
type Z80 struct {
	// the 16 bit registers
	AF, BC, DE, HL, IX, IY uint16

	// their high and low parts
	A, F, B, C, D, E, H, L R8

	// and their alternatives
	AFa, BCa, DEa, HLa uint16

	// the stack pointer and program counter
	SP, PC uint16

	// memory and IO device
	Mem *Memory.RAM
	//IO  io.Device

	// internal flags
	Halted, InterruptEnabled bool

	// inerrupt vector
	I R8

	// flags
	Flag byte

	// Instrctions
	Inst map[z80Instruction]execInstruction
}

func NewCpu(m Memory.RAM) *Z80 {
	z80 := &Z80{
		AF:               0,
		BC:               0,
		DE:               0,
		HL:               0,
		IX:               0,
		IY:               0,
		SP:               0,
		PC:               0,
		Halted:           false,
		InterruptEnabled: false,
		I:                0,
		Flag:             0,
		Mem:              &m,
	}

	z80.Inst = make(map[z80Instruction]execInstruction)

	var zi z80Instruction = 0
	for i := 0; i < 256; i++ {
		z80.registerInstruction(zi, NotImplemented)
		zi++
	}
	z80.registerInstruction(0x00, NOP)
	z80.registerInstruction(0x31, LDSP)

	return z80
}
func (c *Z80) String() string {
	var out bytes.Buffer

	out.WriteString("SP: " + strconv.FormatInt(int64(c.SP), 10))

	out.WriteString(";")
	return out.String()
}

func (c *Z80) registerInstruction(i z80Instruction, fn execInstruction) {
	c.Inst[i] = fn
}

func (c *Z80) Execute() {

	opc := c.Mem.ReadByte(uint16(c.PC))

	exFN := c.Inst[z80Instruction(opc)]

	exFN(c)
}

func (cpu *Z80) IsCary() bool {
	if (cpu.Flag & 0x01) > 0 {
		return true
	} else {
		return false
	}
}

func (cpu *Z80) SetCary() {
	cpu.SetFlag(Carry)
}

func (cpu *Z80) ClearCary() {
	cpu.ClearFlag(Carry)
}

func (cpu *Z80) SetFlag(f byte) {
	cpu.Flag = (cpu.Flag | 1<<f)
}

func (cpu *Z80) ClearFlag(f byte) {
	cpu.Flag = (cpu.Flag | 0<<f)
}
