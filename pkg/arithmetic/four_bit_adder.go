package arithmetic

import (
	"custompro98/circuits/pkg/core"
)

func FourBitAdder(a []core.Bit, b []core.Bit) core.NBitSum {
	result := core.NBitSum{NumBits: 4, S: make([]core.Bit, 4), C: core.BitOff}

	sum4 := HalfAdder(a[3], b[3])
	result.S[3] = sum4.S

	sum3 := FullAdder(a[2], b[2], sum4.C)
	result.S[2] = sum3.S

	sum2 := FullAdder(a[1], b[1], sum3.C)
	result.S[1] = sum2.S

	sum1 := FullAdder(a[0], b[0], sum2.C)
	result.S[0] = sum1.S
	result.C = sum1.C

	return result
}
