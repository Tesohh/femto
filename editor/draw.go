package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (e *Editor) Draw() error {
	// draw Plugins
	for _, p := range e.Plugins {
		switch p.(type) {
		case *DumbPlugin:
			break
		default:
			err := p.Draw(e)
			if err != nil {
				return err
			}
		}
	}

	// draw windows. Draw only orchestrates windows, the drawing itself is done by the windows
	// TODO: Priority can be implemented by simply sorting the windows by priority... 💔🎱
	width, height := e.Screen.Size()
	rights := 0
	lefts := 0
	for _, w := range e.Tab().Windows {
		var err error
		switch w.Alignment {
		// Rights and lefts have priority over Tops and bottoms, so they must be drawn first
		case AlignmentRight:
			err = w.Draw(e, rights, 0, w.Size+rights, height)
			rights += w.Size
		case AlignmentLeft:
			err = w.Draw(e, width-lefts, 0, width-w.Size-lefts, height)
			lefts += w.Size
		}
		if err != nil {
			return err
		}
	}
	tops := 0
	bottoms := 0
	for _, w := range e.Tab().Windows {
		var err error
		// Tops and bottoms (eg. Toolbars) are less important and must be drawn later
		switch w.Alignment {
		case AlignmentTop:
			err = w.Draw(e, rights, tops, width-lefts, w.Size+tops)
			tops += w.Size
		case AlignmentBottom:
			panic("not implemented") // TODO:
		}
		if err != nil {
			return err
		}
	}

	// draw buffer (which will become just a window)
	text, err := e.Buf().Read()
	if err != nil {
		return err
	}

	e.Screen.Clear()

	x := 0
	y := 0

	for _, line := range text {
		for _, char := range line {
			e.Screen.SetContent(x, y, char, nil, tcell.StyleDefault)
			x += runewidth.RuneWidth(char)
		}
		x = 0
		y += 1
	}
	e.Screen.ShowCursor(e.Buf().Pos().X, e.Buf().Pos().Y)

	e.Screen.Show()
	return nil
}
