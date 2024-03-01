package gates

import "custompro98/circuits/pkg/core"

func xnor(a core.Bit, b core.Bit) core.Bit {
	return not(xor(a, b))
}
