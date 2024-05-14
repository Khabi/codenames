package codenames

import "strings"

// Part is any function that can return a string.  This will usually
// be some kind of dictionary lookup that returns a random word.
type Part func() string

// New a generic syntaxically correct slug seperated by dashes.
func New() string {
	return Generate("-", nil, Descriptive, Size, Color, Noun)
}

// Generate creates a string with non-duplicated word parts seperated by a given value.  Specific words may be
// excluded by placing them in the filter slice.
func Generate(separator string, filter []string, parts ...Part) string {
	var out []string
	for _, part := range parts {
	set:
		p := part()
		for _, t := range filter {
			if t == p {
				goto set
			}
		}
		for _, o := range out {
			if o == p {
				goto set
			}
		}
		out = append(out, p)
	}

	formatted := strings.Join(out, separator)
	return formatted
}
