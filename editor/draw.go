package editor

import (
	"cmp"
	"log/slog"
	"slices"
)

func (e *Editor) Draw() error {
	slog.Info("Redrawing...")
	e.Screen.Clear()

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
	windows := append(e.Tab().Windows, e.Windows...)
	windows = slices.DeleteFunc(windows, func(w Window) bool {
		return !w.Shown
	})
	slices.SortFunc(windows, func(a Window, b Window) int {
		return cmp.Compare(a.Priority, b.Priority) * -1
	})

	width, height := e.Screen.Size()
	lefts := 0
	rights := 0
	tops := 0
	bottoms := 0

	for _, w := range windows {
		var err error
		switch w.Alignment {
		case AlignmentLeft:
			err = w.Draw(e, lefts, tops, lefts+w.Size, height-bottoms)
			lefts += w.Size
		case AlignmentRight:
			err = w.Draw(e, width-rights-w.Size, tops, width-rights, height-bottoms)
			rights += w.Size
		case AlignmentTop:
			err = w.Draw(e, lefts, tops, width-rights-rights+1, w.Size) // bandaid
			tops += w.Size
		case AlignmentBottom:
			err = w.Draw(e, lefts, height-bottoms-w.Size, width-rights, height-bottoms)
			bottoms += w.Size
		}
		if err != nil {
			return err
		}
	}

	e.Screen.Show()

	// draw buffer (which will become just a window)
	// text, err := e.Buf().Read()
	// if err != nil {
	// 	return err
	// }
	//
	// x := 0
	// y := 0
	//
	// for _, line := range text {
	// 	for _, char := range line {
	// 		e.Screen.SetContent(x, y, char, nil, tcell.StyleDefault)
	// 		x += runewidth.RuneWidth(char)
	// 	}
	// 	x = 0
	// 	y += 1
	// }

	e.Screen.ShowCursor(e.Buf().Pos().X, e.Buf().Pos().Y)

	// e.Screen.Show()
	return nil
}
