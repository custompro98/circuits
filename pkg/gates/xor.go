package gates

import "custompro98/circuits/pkg/core"

func Xor(a core.Bit, b core.Bit) core.Bit {
	aNandB := Nand(a, b)

	return Nand(Nand(a, aNandB), Nand(b, aNandB))
}
