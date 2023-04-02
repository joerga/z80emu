package z80cpu

import "fmt"

func NotImplemented(c *Z80) {
	fmt.Println("Instruction not implemented")
	c.PC++
}

func NOP(c *Z80) {
	fmt.Println("execute NOP")
	c.PC++
}

func LDSP(c *Z80) {
	fmt.Println("execute LD SP, nn")
	c.PC++
	c.SP = c.Mem.Read16(c.PC)
	c.PC++
	c.PC++
}
