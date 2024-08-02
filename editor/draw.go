package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (e *Editor) Draw() error {
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
