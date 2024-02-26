package gates

import "custompro98/circuits/pkg/core"

func Or(a core.Bit, b core.Bit) core.Bit {
	return Nand(Not(a), Not(b))
}
