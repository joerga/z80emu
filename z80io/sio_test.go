package z80io

import (
	"testing"
)

func TestSIO(t *testing.T) {

	i := NewIO()

	s := NewSIO(0x10, i)

	//t.Errorf("IO: %v", i)

	//t.Errorf("SIO: %v", s)

	t.Errorf("read Chl a: %c", s.readChlAData())
	t.Errorf("read Chl b: %c", s.readChlBData())

	a := i.In(0x10)
	if a != 0x41 {
		t.Errorf("no funktion registred %c", a)
	}

	b := i.In(0x12)
	if b != 0x42 {
		t.Errorf("no funktion registred %c", b)
	}

}
