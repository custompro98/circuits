package gates

import "custompro98/circuits/pkg/core"

func not(a core.Bit) core.Bit {
	return nand(a, a)
}
