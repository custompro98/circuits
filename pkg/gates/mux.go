package gates

import "custompro98/circuits/pkg/core"

func mux(a core.Bit, b core.Bit, s core.Bit) core.Bit {
	return or(and(a, not(s)), and(b, s))
}
