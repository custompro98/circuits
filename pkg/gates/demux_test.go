package gates

import (
	"custompro98/circuits/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemux(t *testing.T) {
	tests := []struct {
		name      string
		a         core.Bit
		s         core.Bit
		expectedA core.Bit
		expectedB core.Bit
	}{
		{
			name:      "On/On yields B",
			a:         core.BitOn,
			s:         core.BitOn,
			expectedA: core.BitOff,
			expectedB: core.BitOn,
		},
		{
			name:      "On/Off yields A",
			a:         core.BitOn,
			s:         core.BitOff,
			expectedA: core.BitOn,
			expectedB: core.BitOff,
		},
		{
			name:      "Off/On yields None",
			a:         core.BitOff,
			s:         core.BitOn,
			expectedA: core.BitOff,
			expectedB: core.BitOff,
		},
		{
			name:      "Off/Off yields None",
			a:         core.BitOff,
			s:         core.BitOff,
			expectedA: core.BitOff,
			expectedB: core.BitOff,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert := assert.New(t)

			a, b := Demux(v.a, v.s)

			assert.Equal(v.expectedA, a, v.name)
			assert.Equal(v.expectedB, b, v.name)
		})
	}
}
