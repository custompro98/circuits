package gates

import "custompro98/circuits/pkg/core"

func And(a core.Bit, b core.Bit) core.Bit {
	return Not(Nand(a, b))
}
