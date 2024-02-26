package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		b        core.Bit
		s        core.Bit
		expected core.Bit
	}{
		{
			name:     "On/On/On yields On",
			a:        core.BitOn,
			b:        core.BitOn,
			s:        core.BitOn,
			expected: core.BitOn,
		},
		{
			name:     "On/Off/On yields Off",
			a:        core.BitOn,
			b:        core.BitOff,
			s:        core.BitOn,
			expected: core.BitOff,
		},
		{
			name:     "Off/On/On yields On",
			a:        core.BitOff,
			b:        core.BitOn,
			s:        core.BitOn,
			expected: core.BitOn,
		},
		{
			name:     "Off/Off/On yields Off",
			a:        core.BitOff,
			b:        core.BitOff,
			s:        core.BitOn,
			expected: core.BitOff,
		},
		{
			name:     "On/On/Off yields On",
			a:        core.BitOn,
			b:        core.BitOn,
			s:        core.BitOff,
			expected: core.BitOn,
		},
		{
			name:     "On/Off/Off yields On",
			a:        core.BitOn,
			b:        core.BitOff,
			s:        core.BitOff,
			expected: core.BitOn,
		},
		{
			name:     "Off/On/Off yields Off",
			a:        core.BitOff,
			b:        core.BitOn,
			s:        core.BitOff,
			expected: core.BitOff,
		},
		{
			name:     "Off/Off/Off yields Off",
			a:        core.BitOff,
			b:        core.BitOff,
			s:        core.BitOff,
			expected: core.BitOff,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := Mux(v.a, v.b, v.s)

			assert.Equal(result, v.expected, v.name)
		})
	}
}
