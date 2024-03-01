package gates

import "custompro98/circuits/pkg/core"

func xor(a core.Bit, b core.Bit) core.Bit {
	aNandB := nand(a, b)

	return nand(nand(a, aNandB), nand(b, aNandB))
}
