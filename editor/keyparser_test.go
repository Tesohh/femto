package editor

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSplitHumanKeySequence(t *testing.T) {
	assert.DeepEqual(t, []string{"x", "<C-a>"}, SplitHumanKeySequence("x<C-a>"))
	assert.DeepEqual(t, []string{"x", "<F10>", "y", "z"}, SplitHumanKeySequence("x<F10>yz"))
}
