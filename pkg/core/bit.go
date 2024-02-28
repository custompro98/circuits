package core

type Bit int

const (
	BitOff Bit = iota
	BitOn
)

type Sum struct {
	S Bit
	C Bit
}

type NBitSum struct {
	NumBits int
	S []Bit
	C Bit
}
