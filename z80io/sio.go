package z80io

import "fmt"

// register of one SIO channel
type sioChannel struct {
	wr0, wr1, wr2, wr3, wr4, wr5, wr6, wr7 byte
	rr0, rr1, rr2                          byte
}

// 2 SIO Channel and a reference to IO dispatcher
type SIO struct {
	//	ChlAData byte
	//	ChlACtrl byte
	//	ChlBData byte
	//	ChlBCtrl byte

	ChannelA sioChannel
	ChannelB sioChannel
	i        *Z80IO
}

func NewSIO(baseaddr IOAddr, io *Z80IO) *SIO {
	s := &SIO{
		i: io,
		ChannelA: sioChannel{
			wr0: 0,
			wr1: 0,
			wr2: 0,
			wr3: 0,
			wr4: 0,
			wr5: 0,
			wr6: 0,
			wr7: 0,
			rr0: 0,
			rr1: 0,
			rr2: 0,
		},
		ChannelB: sioChannel{
			wr0: 0,
			wr1: 0,
			wr2: 0,
			wr3: 0,
			wr4: 0,
			wr5: 0,
			wr6: 0,
			wr7: 0,
			rr0: 0,
			rr1: 0,
			rr2: 0,
		},
	}

	s.i.registerOutputDevice(baseaddr, s.writeChlAData)
	s.i.registerOutputDevice(baseaddr+1, s.writeChlACTRL)
	s.i.registerOutputDevice(baseaddr+2, s.writeChlBData)
	s.i.registerOutputDevice(baseaddr+3, s.writeChlBCTRL)

	s.i.registerInputDevice(baseaddr, s.readChlAData)
	s.i.registerInputDevice(baseaddr+2, s.readChlBData)

	return s
}

func (s *SIO) writeChlACTRL(val byte) {

}

func (s *SIO) writeChlAData(val byte) {
	fmt.Print("Write auf auf SIO Channel A")

}
func (s *SIO) writeChlBCTRL(val byte) {

}
func (s *SIO) writeChlBData(val byte) {
	fmt.Print("Write auf auf SIO Channel B")
}

func (s *SIO) readChlAData() byte {
	return 0x41

}
func (s *SIO) readChlBData() byte {
	return 0x42
}
