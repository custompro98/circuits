package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNand(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		b        core.Bit
		expected core.Bit
	}{
		{
			name:     "On/On yields Off",
			a:        core.BitOn,
			b:        core.BitOn,
			expected: core.BitOff,
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
			name:     "Off/Off yields On",
			a:        core.BitOff,
			b:        core.BitOff,
			expected: core.BitOn,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := Nand(v.a, v.b)

			assert.Equal(result, v.expected, v.name)
		})
	}
}
