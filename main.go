package main

import (
	"fmt"
	"io"
	"os"
	"z80emu/Memory"
	"z80emu/repl"
	"z80emu/z80cpu"
	"z80emu/z80io"
)

func main() {
	f, err := os.Open("sio.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	inp, _ := io.ReadAll(f)

	ram := Memory.NewRAM()
	copy(ram.Data, inp)

	i := z80io.NewIO()

	z80io.NewSIO(0xF0, i)

	cpu := z80cpu.NewCpu(*ram)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout, cpu)

	//fmt.Printf("%s", hex.Dump(ram.Data))

}
