package arithmetic

import (
	"custompro98/circuits/pkg/core"
)

func NBitAdder(a []core.Bit, b []core.Bit) core.NBitSum {
	numBits := max(len(a), len(b))
	result := core.NBitSum{NumBits: numBits, S: make([]core.Bit, numBits), C: core.BitOff}

	sum := core.Sum{}
	for i := result.NumBits - 1; i >= 0; i-- {
		sum = FullAdder(a[i], b[i], sum.C)

		result.S[i] = sum.S

		if i == 0 {
			result.C = sum.C
		}
	}

	return result
}
