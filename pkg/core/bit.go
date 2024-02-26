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
