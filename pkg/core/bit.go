package core

type Bit int

const (
	BitOff     Bit = iota
	BitOn
	BitInvalid Bit = -1
)

type Sum struct {
	S Bit
	C Bit
}

type NBitSum struct {
	NumBits int
	S       []Bit
	C       Bit
}
