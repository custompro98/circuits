package gates

import "custompro98/circuits/pkg/core"

func nor(a core.Bit, b core.Bit) core.Bit {
	return not(nand(not(a), not(b)))
}
