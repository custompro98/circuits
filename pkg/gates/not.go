package gates

import "custompro98/circuits/pkg/core"

func Not(a core.Bit) core.Bit {
	return Nand(a, a)
}
