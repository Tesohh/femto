package humankey

import "github.com/gdamore/tcell/v2"

type InternalKey struct {
	Key     tcell.Key
	Rune    rune
	ModMask tcell.ModMask
}

func (i InternalKey) Matches(event tcell.EventKey) bool {
	return i.Key == event.Key() && i.Rune == event.Rune() && i.ModMask == event.Modifiers()
}
