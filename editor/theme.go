package editor

import "github.com/gdamore/tcell/v2"

type Theme struct {
	Name string

	Default tcell.Style
	Borders tcell.Color

	Error tcell.Style

	Red       tcell.Color
	Yellow    tcell.Color
	Pink      tcell.Color
	Blue      tcell.Color
	LightBlue tcell.Color
	Purple    tcell.Color

	NormalModeAccent tcell.Color
	InsertModeAccent tcell.Color
}
