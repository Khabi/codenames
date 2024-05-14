package codenames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dupes  = []string{"GDO", "GDO", "DHD"}
	filter = []string{"goa'uld", "tau'ri"}
)

func name() string { return "stargate" }
func abbr() string { return "sg" }
func num() string  { return "1" }
func dupe() string {
	x := dupes[0]
	dupes = append(dupes[:0], dupes[1:]...)
	return x
}
func filtered() string {
	x := filter[0]
	filter = append(filter[:0], filter[1:]...)
	return x
}

func TestNew(t *testing.T) {
	assert.NotEqual(t, "", New())
}

func TestGenerate(t *testing.T) {
	var tests = []struct {
		parts     []Part
		seperator string
		filter    []string
		expected  string
	}{
		{
			parts:     []Part{name, abbr, num},
			seperator: "-",
			filter:    nil,
			expected:  "stargate-sg-1",
		},
		{
			parts:     []Part{dupe, dupe},
			seperator: "-",
			filter:    nil,
			expected:  "GDO-DHD",
		},
		{
			parts:     []Part{name, filtered},
			seperator: " ",
			filter:    []string{"nomatch", "goa'uld"},
			expected:  "stargate tau'ri",
		},
	}

	for _, tt := range tests {
		res := Generate(tt.seperator, tt.filter, tt.parts...)
		assert.Equal(t, tt.expected, res)
	}
}
