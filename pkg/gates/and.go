package gates

import "custompro98/circuits/pkg/core"

func and(a core.Bit, b core.Bit) core.Bit {
	return not(nand(a, b))
}
