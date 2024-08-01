package humankey

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

func PrettyPrintInternalKey(key InternalKey) string {
	if key.Key >= tcell.KeyCtrlA && key.Key <= tcell.KeyCtrlUnderscore {
		return tcell.KeyNames[key.Key]
	}

	s := ""
	if key.Key != tcell.KeyRune {
		s += tcell.KeyNames[key.Key]
	} else {
		s += string(key.Rune)
	}

	if key.ModMask&tcell.ModCtrl != 0 {
		s = "ctrl+" + s
	}
	if key.ModMask&tcell.ModShift != 0 {
		s = "shift+" + s
	}
	if key.ModMask&tcell.ModAlt != 0 {
		s = "alt+" + s
	}
	if key.ModMask&tcell.ModMeta != 0 {
		s = "meta+" + s
	}

	return s
}

func PrettyPrintSequence(seq []InternalKey) string {
	s := ""
	for _, v := range seq {
		s += PrettyPrintInternalKey(v) + " "
	}

	s = strings.TrimSuffix(s, " ")
	return s
}
