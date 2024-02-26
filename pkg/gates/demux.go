package gates

import "custompro98/circuits/pkg/core"

func Demux(a core.Bit, s core.Bit) (core.Bit, core.Bit) {
	outputA := Not(Nand(a, Not(s)))
	outputB := Not(Nand(a, s))

	return outputA, outputB
}
