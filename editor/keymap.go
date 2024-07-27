package editor

import (
	"github.com/gdamore/tcell/v2"
)

type Keymap map[string]map[tcell.Key]string

// type Key struct {
// 	Key  tcell.Key
// 	Rune rune
// }

var defaultKeymap = Keymap{
	ModeNormal: {
		tcell.KeyCtrlC: "quit",
	},
}
