package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"custompro98/circuits/pkg/gates"
)

func HalfAdder(a core.Bit, b core.Bit) core.Sum {
	return core.Sum{
		S: gates.Xor(a, b),
		C: gates.And(a, b),
	}
}
