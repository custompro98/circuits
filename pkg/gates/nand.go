package gates

import "custompro98/circuits/pkg/core"

func nand(a core.Bit, b core.Bit) core.Bit {
	if a == core.BitOn && b == core.BitOn {
		return core.BitOff
	}

	return core.BitOn
}
