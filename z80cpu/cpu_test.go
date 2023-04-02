package z80cpu

import (
	"testing"
	"z80emu/Memory"
)

func TestFlags(t *testing.T) {

	ram := Memory.NewRAM()
	z := NewCpu(*ram)

	if z.IsCary() {
		t.Errorf("Error Caryflag not reset")
	}

	z.SetCary()
	if !z.IsCary() {
		t.Errorf("Error Caryflag not set")
	}

	z.ClearCary()
	if !z.IsCary() {
		t.Errorf("Error Caryflag not cleared")
	}

}
