package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourBitAdder(t *testing.T) {
	tests := []struct {
		name     string
		a        []core.Bit
		b        []core.Bit
		expected core.NBitSum
	}{
		{
			name: "Zero plus Zero yields Zero/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "Zero plus One yields One/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
				C:       core.BitOff,
			},
		},
		{
			name: "Zero plus Two yields Two/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOn, core.BitOff},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOn, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "Zero plus Three yields Three/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOn, core.BitOn},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOn, core.BitOn},
				C:       core.BitOff,
			},
		},
		{
			name: "Zero plus Four yields Four/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOff, core.BitOn, core.BitOff, core.BitOff},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOn, core.BitOff, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "Zero plus Sixteen yields Sixteen/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
			b:    []core.Bit{core.BitOn, core.BitOff, core.BitOff, core.BitOff},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOn, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "One plus One yields Two/NoCarry",
			a:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOn, core.BitOff},
				C:       core.BitOff,
			},
		},
		{
			name: "Thirty-one plus One yields Zero/Carry",
			a:    []core.Bit{core.BitOn, core.BitOn, core.BitOn, core.BitOn},
			b:    []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOn},
			expected: core.NBitSum{
				NumBits: 4,
				S:       []core.Bit{core.BitOff, core.BitOff, core.BitOff, core.BitOff},
				C:       core.BitOn,
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := FourBitAdder(v.a, v.b)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
