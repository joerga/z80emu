package repl

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"z80emu/z80cpu"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, c *z80cpu.Z80) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := strings.Split(line, " ")

		s := make([]byte, 128)
		copy(s, []byte(line))

		fmt.Printf("%v\n", l)

		switch s[0] {
		case 'd':
			dumpMemory(0x0000, c)
		case 'r':
			fmt.Println(c.String())
		case 's':
			cpuStep(c)
		case 'q':
			os.Exit(0)

		}

		io.WriteString(out, line)
		io.WriteString(out, "\n")
	}
}

func dumpMemory(start uint16, c *z80cpu.Z80) {

	s := make([]byte, 256)
	copy(s, c.Mem.Data[0:256])

	fmt.Printf("%s", hex.Dump(s))
}

func cpuStep(c *z80cpu.Z80) {
	c.Execute()
}
