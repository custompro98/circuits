package arithmetic

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullAdder(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		b        core.Bit
		c        core.Bit
		expected core.Sum
	}{
		{
			name:     "On/On/On yields On/Carry",
			a:        core.BitOn,
			b:        core.BitOn,
			c:        core.BitOn,
			expected: core.Sum{S: core.BitOn, C: core.BitOn},
		},
		{
			name:     "On/On/Off yields Off/Carry",
			a:        core.BitOn,
			b:        core.BitOn,
			c:        core.BitOff,
			expected: core.Sum{S: core.BitOff, C: core.BitOn},
		},
		{
			name:     "On/Off/On yields Off/Carry",
			a:        core.BitOn,
			b:        core.BitOff,
			c:        core.BitOn,
			expected: core.Sum{S: core.BitOff, C: core.BitOn},
		},
		{
			name:     "On/Off/Off yields On/NoCarry",
			a:        core.BitOn,
			b:        core.BitOff,
			c:        core.BitOff,
			expected: core.Sum{S: core.BitOn, C: core.BitOff},
		},
		{
			name:     "Off/On/On yields Off/Carry",
			a:        core.BitOff,
			b:        core.BitOn,
			c:        core.BitOn,
			expected: core.Sum{S: core.BitOff, C: core.BitOn},
		},
		{
			name:     "Off/On/Off yields On/NoCarry",
			a:        core.BitOff,
			b:        core.BitOn,
			c:        core.BitOff,
			expected: core.Sum{S: core.BitOn, C: core.BitOff},
		},
		{
			name:     "Off/Off/On yields On/NoCarry",
			a:        core.BitOff,
			b:        core.BitOff,
			c:        core.BitOn,
			expected: core.Sum{S: core.BitOn, C: core.BitOff},
		},
		{
			name:     "Off/Off/Off yields Off",
			a:        core.BitOff,
			b:        core.BitOff,
			c:        core.BitOff,
			expected: core.Sum{S: core.BitOff, C: core.BitOff},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := FullAdder(v.a, v.b, v.c)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
