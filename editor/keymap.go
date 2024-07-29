package editor

import (
	"github.com/gdamore/tcell/v2"
)

type Keymap map[Mode]map[tcell.Key]string

var defaultKeymap = Keymap{
	ModeNormal: {
		tcell.KeyCtrlC: "quit",
	},
}
