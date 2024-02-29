package arithmetic

import (
	"custompro98/circuits/pkg/core"
)

func NBitAdder(a []core.Bit, b []core.Bit) core.NBitSum {
	numBits := max(len(a), len(b))
	result := core.NBitSum{NumBits: numBits, S: make([]core.Bit, numBits), C: core.BitOff}

	aprime := normalizeInput(numBits, a)
	bprime := normalizeInput(numBits, b)

	sum := core.Sum{}
	for i := result.NumBits - 1; i >= 0; i-- {
		sum = FullAdder(aprime[i], bprime[i], sum.C)

		result.S[i] = sum.S

		if i == 0 {
			result.C = sum.C
		}
	}

	return result
}

func normalizeInput(numBits int, input []core.Bit) []core.Bit {
	prime := make([]core.Bit, numBits)

	// Iterate over the input from LSB to MSB.
	// Backfill the difference from numBits to
	// input's MSB position with core.BitOff
	for i := 0; i < len(input); i++ {
		target := numBits - 1 - i
		source := len(input) - 1 - i
		prime[target] = input[source]

	}

	return prime
}
