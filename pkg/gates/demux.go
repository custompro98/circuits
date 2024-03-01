package gates

import "custompro98/circuits/pkg/core"

func demux(a core.Bit, s core.Bit) (core.Bit, core.Bit) {
	outputA := not(nand(a, not(s)))
	outputB := not(nand(a, s))

	return outputA, outputB
}
