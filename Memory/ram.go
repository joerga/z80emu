package Memory

import (
	"log"
)

// RAM represents the RAM in the Z80
type RAM struct {
	Data []byte
}

const ramSize = 64 * 1024 // 64kbyte

// NewRAM makes a new RAM object with size 64k and populates with the initial data supplied
func NewRAM() *RAM {
	ram := RAM{
		Data: make([]byte, ramSize),
	}
	return &ram
}

func (r *RAM) ReadByte(addr uint16) byte {
	return r.Data[addr]
}

func (r *RAM) Write(addr uint16, data byte) {
	if addr > ramSize-1 {
		log.Panic("[RAM] Tried to write outside RAM")
	}
	r.Data[addr] = data
}

func (r *RAM) Read16(addr uint16) uint16 {
	return uint16(r.Data[addr+1])<<8 | uint16(r.Data[addr])
}
