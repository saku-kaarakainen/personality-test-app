package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnformat(t *testing.T) {

	var tests = []struct {
		format, input, want string
	}{
		{"<p>%s</p>", "<p>Hello</p>", "Hello"},
		{"q[%s]", "q[3]", "3"},
		{"{ 'key': '%s' }", "{ 'key': 'value' }", "value"},
		{"test:%sThe rest of this is ignored.%s", "test:3T", "3"},
		// There is no error handling for indexing, so this might fail in special cases, like if %s is at end.
	}

	for i, tt := range tests {
		testname := fmt.Sprintf(
			"%d ('%s', '%s', '%s')",
			i, tt.format, tt.input, tt.want)

		t.Run(testname, func(t *testing.T) {
			actual, _ := Unformat(tt.format, tt.input)
			assert.Equal(t, tt.want, actual)
		})
	}
}
