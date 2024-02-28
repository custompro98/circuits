package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"custompro98/circuits/pkg/gates"
)

func FullAdder(a core.Bit, b core.Bit, c core.Bit) core.Sum {
	aPlusB := HalfAdder(a, b)
	sumPlusCarry := HalfAdder(aPlusB.S, c)

	return core.Sum{
		S: sumPlusCarry.S,
		C: gates.Or(aPlusB.C, sumPlusCarry.C),
	}
}
