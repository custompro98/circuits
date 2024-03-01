package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOr(t *testing.T) {
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
			name:     "On/Off yields On",
			a:        core.BitOn,
			b:        core.BitOff,
			expected: core.BitOn,
		},
		{
			name:     "Off/On yields On",
			a:        core.BitOff,
			b:        core.BitOn,
			expected: core.BitOn,
		},
		{
			name:     "Off/Off yields Off",
			a:        core.BitOff,
			b:        core.BitOff,
			expected: core.BitOff,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := or(v.a, v.b)

			assert.Equal(v.expected, result, v.name)
		})
	}
}
