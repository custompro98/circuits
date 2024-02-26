package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name     string
		a        core.Bit
		expected core.Bit
	}{
		{
			name:     "On yields Off",
			a:        core.BitOn,
			expected: core.BitOff,
		},
		{
			name:     "Off yields On",
			a:        core.BitOff,
			expected: core.BitOn,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			result := Not(v.a)

			assert.Equal(result, v.expected, v.name)
		})
	}
}
