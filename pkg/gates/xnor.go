package gates

import "custompro98/circuits/pkg/core"

func Xnor(a core.Bit, b core.Bit) core.Bit {
	return Not(Xor(a, b))
}
