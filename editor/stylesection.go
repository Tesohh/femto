package editor

import "github.com/gdamore/tcell/v2"

type StyleSection struct {
	Y      int
	StartX int
	EndX   int
	Style  tcell.Style
}
