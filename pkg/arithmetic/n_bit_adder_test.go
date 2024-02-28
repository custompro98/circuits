package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNBitAdder(t *testing.T) {
	tests := []struct {
		name     string
		a        []core.Bit
		b        []core.Bit
		expected core.NBitSum
	}{
		{
			name: "Zero plus Zero (4-bit) yields Zero/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "Thirty-one plus One (4-bit) yields Zero/Carry",
			a:    []core.Bit{core.BitOn, core.BitOn, core.BitOn, core.BitOn},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOn,
			},
		},
		{
			name: "One plus One (1-bit) yields Zero/Carry",
			a:    []core.Bit{core.BitOn},
			b:    []core.Bit{core.BitOn},
			expected: core.NBitSum{
				NumBits: 1,
				S:       []core.Bit{core.BitOff},
				C:       core.BitOn,
			},
		},
		{
			name: "Two-hundred Fifty-Five plus One (8-bit) yields Zero/Carry",
			a:    []core.Bit{core.BitOn, core.BitOn, core.BitOn, core.BitOn, core.BitOn, core.BitOn, core.BitOn, core.BitOn},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			expected: core.NBitSum{
				NumBits: 8,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOn,
			},
		},
		{
			name: "Five plus Five (3-bit) yields Four/Carry",
			a:    []core.Bit{core.BitOn, core.BitOn, core.BitOn},
			b:    []core.Bit{core.BitOn, core.BitOn, core.BitOn},
			expected: core.NBitSum{
				NumBits: 3,
				S:       []core.Bit{core.BitOn, core.BitOn, core.BitOff},
				C:       core.BitOn,
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := NBitAdder(v.a, v.b)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
