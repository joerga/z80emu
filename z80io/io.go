package z80io

// Device defines the interface for communicating with an IO device using IN and OUT operations
type Device interface {
	// DeviceWrite writes a single byte to the specified port
	DeviceOut(val byte)

	// DeviceRead reads a single byte from the specified port
	DeviceIn() byte
}
type IOAddr byte

type (
	outFn func(port byte)
	inFn  func() byte
)

type Z80IO struct {
	outFns map[IOAddr]outFn
	inFns  map[IOAddr]inFn

	InterruptPending bool
	InterruptVector  byte
}

func NewIO() *Z80IO {
	io := &Z80IO{
		InterruptPending: false,
		InterruptVector:  0,
	}

	io.outFns = make(map[IOAddr]outFn)
	io.inFns = make(map[IOAddr]inFn)

	return io
}

func (io *Z80IO) registerOutputDevice(addr IOAddr, fn outFn) {
	io.outFns[addr] = fn
}

func (io *Z80IO) registerInputDevice(addr IOAddr, fn inFn) {
	io.inFns[addr] = fn
}

func (io *Z80IO) In(addr IOAddr) byte {
	inFn := io.inFns[addr]
	if inFn == nil {
		panic("no input FN defined")
	}
	return inFn()
}
