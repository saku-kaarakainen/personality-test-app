package result

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertScoreToFlag(t *testing.T) {

	var tests = []struct {
		score        [2]int32
		expectedFlag int32
	}{
		{[2]int32{-2, -2}, 0},
		{[2]int32{2, -2}, 1},
		{[2]int32{-2, 2}, 2},
		{[2]int32{2, 2}, 3},
		// There is no error handling for indexing, so this might fail in special cases, like if %s is at end.
	}

	for i, tt := range tests {
		testname := fmt.Sprintf(
			"%d (for score:'%v', expected:'%d')",
			i, tt.score, tt.expectedFlag)

		t.Run(testname, func(t *testing.T) {
			actualFlag := convertScoreToFlag(tt.score)
			assert.Equal(t, tt.expectedFlag, actualFlag)
		})
	}
}
