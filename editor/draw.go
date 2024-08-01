package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (e *Editor) Draw() error {
	buf := e.tab().Buffer
	text, err := buf.Read()
	if err != nil {
		return err
	}

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
	e.Screen.ShowCursor(buf.Pos().X, buf.Pos().Y)

	e.Screen.Show()
	return nil
}
