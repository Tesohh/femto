package humankey

import "github.com/gdamore/tcell/v2"

type InternalKey struct {
	Key     tcell.Key
	Rune    rune
	ModMask tcell.ModMask
}
