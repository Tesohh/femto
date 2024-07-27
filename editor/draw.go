package editor

import "github.com/gdamore/tcell/v2"

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
			x += 1
		}
		x = 0
		y += 1
	}
	e.Screen.ShowCursor(x, y)

	e.Screen.Show()
	return nil
}
