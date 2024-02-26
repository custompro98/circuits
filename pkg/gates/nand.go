package gates

import "custompro98/circuits/pkg/core"

func Nand(a core.Bit, b core.Bit) core.Bit {
	if a == core.BitOn && b == core.BitOn {
		return core.BitOff
	}

	return core.BitOn
}
