package humankey

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"gotest.tools/v3/assert"
)

func TestInternalKeymapGetMatches(t *testing.T) {
	humanSampleKeymap := HumanKeymap{
		"normal": {
			"c i w": "change_in_word",
			"c i p": "change_in_paragraph",
			"c c":   "change_line",
		},
	}
	keymap, err := humanSampleKeymap.ToInternal()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("match c to sample should be len(3)", func(t *testing.T) {
		currentSequence := []InternalKey{{Rune: 'c', Key: tcell.KeyRune}}
		matches := keymap.GetMatches("normal", currentSequence)
		assert.Equal(t, len(matches), 3)
	})

	t.Run("match ci to sample should be len(2)", func(t *testing.T) {
		currentSequence := []InternalKey{{Rune: 'c', Key: tcell.KeyRune}, {Rune: 'i', Key: tcell.KeyRune}}
		matches := keymap.GetMatches("normal", currentSequence)
		assert.Equal(t, len(matches), 2)
	})

	t.Run("match ciw to sample should be len(1)", func(t *testing.T) {
		currentSequence := []InternalKey{{Rune: 'c', Key: tcell.KeyRune}, {Rune: 'i', Key: tcell.KeyRune}, {Rune: 'w', Key: tcell.KeyRune}}
		matches := keymap.GetMatches("normal", currentSequence)
		assert.Equal(t, len(matches), 1)
	})
}
