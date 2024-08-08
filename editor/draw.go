package editor

import (
	"cmp"
	"log/slog"
	"reflect"
	"slices"

	"github.com/gdamore/tcell/v2"
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
		isFocused := reflect.DeepEqual(*e.Win(), w)
		switch w.Alignment {
		case AlignmentLeft:
			err = w.Draw(e, lefts, tops, lefts+w.Size, height-bottoms, isFocused)
			lefts += w.Size
		case AlignmentRight:
			err = w.Draw(e, width-rights-w.Size, tops, width-rights, height-bottoms, isFocused)
			rights += w.Size
		case AlignmentTop:
			err = w.Draw(e, lefts, tops, width-rights-rights+1, w.Size, isFocused) // bandaid
			tops += w.Size
		case AlignmentBottom:
			err = w.Draw(e, lefts, height-bottoms-w.Size, width-rights, height-bottoms, isFocused)
			bottoms += w.Size
		}
		if err != nil {
			return err
		}
	}

	// render the center window (presumably the buffer)
	tempcount := 0
	for _, w := range windows {
		if w.Alignment != AlignmentCenter {
			continue
		}
		if tempcount > 1 {
			panic("more than 1 centered window. currently impossible")
		}

		isFocused := reflect.DeepEqual(*e.Win(), w)
		err := w.Draw(e, lefts, tops, width-rights, height-bottoms, isFocused)
		if err != nil {
			return err
		}

		tempcount += 1
	}

	switch e.Win().Mode {
	case "normal":
		e.Screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	case "insert":
		e.Screen.SetCursorStyle(tcell.CursorStyleSteadyBar)
	}
	e.Screen.Show()
	return nil
}
