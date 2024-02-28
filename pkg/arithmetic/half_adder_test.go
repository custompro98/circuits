package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalfAdder(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		b        core.Bit
		expected core.Sum
	}{
		{
			name:     "On/On yields Off/Carry",
			a:        core.BitOn,
			b:        core.BitOn,
			expected: core.Sum{S: core.BitOff, C: core.BitOn},
		},
		{
			name:     "On/Off yields On/NoCarry",
			a:        core.BitOn,
			b:        core.BitOff,
			expected: core.Sum{S: core.BitOn, C: core.BitOff},
		},
		{
			name:     "Off/On yields On/NoCarry",
			a:        core.BitOff,
			b:        core.BitOn,
			expected: core.Sum{S: core.BitOn, C: core.BitOff},
		},
		{
			name:     "Off/Off yields Off/NoCarry",
			a:        core.BitOff,
			b:        core.BitOff,
			expected: core.Sum{S: core.BitOff, C: core.BitOff},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := HalfAdder(v.a, v.b)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
