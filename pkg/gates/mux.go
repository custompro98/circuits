package gates

import "custompro98/circuits/pkg/core"

func Mux(a core.Bit, b core.Bit, s core.Bit) core.Bit {
	return Or(And(a, Not(s)), And(b, s))
}
