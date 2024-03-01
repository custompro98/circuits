package gates

import "custompro98/circuits/pkg/core"

func or(a core.Bit, b core.Bit) core.Bit {
	return nand(not(a), not(b))
}
