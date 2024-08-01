package humankey

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {
	t.Run("just a rune should return just a rune", func(t *testing.T) {
		x, _ := Parse("a")
		assert.Equal(t, x, InternalKey{Rune: 'a', Key: tcell.KeyRune})
	})

	t.Run("capital rune should not become rune+modshift", func(t *testing.T) {
		x, _ := Parse("A")
		assert.Equal(t, x, InternalKey{Rune: 'A', Key: tcell.KeyRune})
	})

	t.Run("special keys should become a key", func(t *testing.T) {
		x, _ := Parse("F11")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyF11})
	})

	t.Run("\"space\" should become ' '", func(t *testing.T) {
		x, _ := Parse("space")
		assert.Equal(t, x, InternalKey{Rune: ' ', Key: tcell.KeyRune})
	})

	t.Run("ctrl+letter combo should return special tcell.KeyCtrl*", func(t *testing.T) {
		x, _ := Parse("ctrl+a")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyCtrlA, Rune: 'a', ModMask: tcell.ModCtrl})
	})

	t.Run("ctrl+specialchar combo should return rune+modctrl", func(t *testing.T) {
		x, _ := Parse("ctrl+@")
		assert.Equal(t, x, InternalKey{Rune: '@', Key: tcell.KeyRune, ModMask: tcell.ModCtrl})
	})

	t.Run("multiple modifiers should be masked together", func(t *testing.T) {
		x, _ := Parse("ctrl+alt+A")
		assert.Equal(t, x, InternalKey{Rune: 'A', Key: tcell.KeyRune, ModMask: tcell.ModCtrl | tcell.ModAlt})
	})

	t.Run("ctrl+specialkey should return key+modctrl", func(t *testing.T) {
		x, _ := Parse("ctrl+F12")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyF12, ModMask: tcell.ModCtrl})
	})
}
