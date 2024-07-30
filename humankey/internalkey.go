package humankey

import "github.com/gdamore/tcell/v2"

type InternalKey struct {
	Key     tcell.Key
	Rune    rune
	ModMask tcell.ModMask
}

func (i InternalKey) Matches(event tcell.EventKey) bool {
	keyMatches := i.Key == event.Key()
	runeMatches := i.Rune == event.Rune() || event.Key() != tcell.KeyRune
	modMaskMatches := i.ModMask == event.Modifiers()

	return keyMatches && runeMatches && modMaskMatches
}

func (i InternalKey) MatchesInternal(key InternalKey) bool {
	keyMatches := i.Key == key.Key
	runeMatches := i.Rune == key.Rune || key.Key != tcell.KeyRune
	modMaskMatches := i.ModMask == key.ModMask

	return keyMatches && runeMatches && modMaskMatches
}
