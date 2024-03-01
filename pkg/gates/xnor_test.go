package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXnor(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		b        core.Bit
		expected core.Bit
	}{
		{
			name:     "On/On yields On",
			a:        core.BitOn,
			b:        core.BitOn,
			expected: core.BitOn,
		},
		{
			name:     "On/Off yields Off",
			a:        core.BitOn,
			b:        core.BitOff,
			expected: core.BitOff,
		},
		{
			name:     "Off/On yields Off",
			a:        core.BitOff,
			b:        core.BitOn,
			expected: core.BitOff,
		},
		{
			name:     "Off/Off yields On",
			a:        core.BitOff,
			b:        core.BitOff,
			expected: core.BitOn,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := xnor(v.a, v.b)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
