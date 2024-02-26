package gates

import "custompro98/circuits/pkg/core"

func Nor(a core.Bit, b core.Bit) core.Bit {
	return Not(Nand(Not(a), Not(b)))
}
