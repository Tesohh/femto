package humankey

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {
	t.Run("just a rune should return just a rune", func(t *testing.T) {
		x, _ := Parse("a")
		assert.Equal(t, x, InternalKey{Rune: 'a'})
	})

	t.Run("capital rune should become rune+modshift", func(t *testing.T) {
		x, _ := Parse("A")
		assert.Equal(t, x, InternalKey{Rune: 'a', ModMask: tcell.ModShift})
	})

	t.Run("special keys should become a key", func(t *testing.T) {
		x, _ := Parse("F11")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyF11})
	})

	t.Run("\"space\" should become ' '", func(t *testing.T) {
		x, _ := Parse("space")
		assert.Equal(t, x, InternalKey{Rune: ' '})
	})

	t.Run("ctrl+letter combo should return special tcell.KeyCtrl*", func(t *testing.T) {
		x, _ := Parse("ctrl+a")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyCtrlA})
	})

	t.Run("ctrl+specialchar combo should return rune+modctrl", func(t *testing.T) {
		x, _ := Parse("ctrl+@")
		assert.Equal(t, x, InternalKey{Rune: '@', ModMask: tcell.ModCtrl})
	})

	t.Run("ctrl+Capital and ctrl+shift+lower should be the same", func(t *testing.T) {
		x, _ := Parse("ctrl+A")
		y, _ := Parse("ctrl+shift+a")
		assert.Equal(t, x, InternalKey{Rune: 'a', ModMask: tcell.ModCtrl | tcell.ModShift})
		assert.Equal(t, x, y)
	})

	t.Run("multiple modifiers should be masked together and capital letter becomes shift", func(t *testing.T) {
		x, _ := Parse("ctrl+alt+A")
		assert.Equal(t, x, InternalKey{Rune: 'a', ModMask: tcell.ModCtrl | tcell.ModShift | tcell.ModAlt})
	})

	t.Run("ctrl+specialkey should return key+modctrl", func(t *testing.T) {
		x, _ := Parse("ctrl+F12")
		assert.Equal(t, x, InternalKey{Key: tcell.KeyF12, ModMask: tcell.ModCtrl})
	})
}
