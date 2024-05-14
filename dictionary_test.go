package codenames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		source   []string
		function Part
	}{
		{
			source:   age,
			function: Age,
		},
		{
			source:   color,
			function: Color,
		},
		{
			source:   descriptive,
			function: Descriptive,
		},
		{
			source:   futurama,
			function: Futurama,
		},
		{
			source:   noun,
			function: Noun,
		},
		{
			source:   planet,
			function: Planet,
		},
		{
			source:   shape,
			function: Shape,
		},
		{
			source:   size,
			function: Size,
		},
	}

	for _, tt := range tests {
		word := tt.function()
		exists := false
		for _, w := range tt.source {
			if w == word {
				exists = true
			}
		}

		assert.True(t, exists)
	}
}
