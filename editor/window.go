package editor

import (
	"log/slog"

	"github.com/Tesohh/femto/buffer"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (e *Editor) RegisterWindow(w Window) {
	e.Tab().Windows = append(e.Tab().Windows, w)
}

type Alignment uint8

const (
	AlignmentLeft   Alignment = 0
	AlignmentRight  Alignment = 1
	AlignmentTop    Alignment = 2
	AlignmentBottom Alignment = 3
)

type WindowFlags uint8

const (
	WindowFlagReadonly    WindowFlags = 1
	WindowFlagInteractive WindowFlags = 2
	WindowFlagHasBorder   WindowFlags = 4
)

type Window struct {
	Alignment Alignment
	Size      int
	Priority  int

	Shown   bool
	Focused bool
	Flags   WindowFlags

	Content buffer.Buffer
	// TODO: ColorSections
}

func (w *Window) Draw(e *Editor, startX int, startY int, boundX int, boundY int) error {
	text, err := w.Content.Read()
	if err != nil {
		return err
	}

	slog.Info("SLOG THIS SHIT")

	x := 0
	y := 0

	for _, line := range text {
		for _, char := range line {
			if x+startX >= boundX {
				continue
			}
			e.Screen.SetContent(x+startX, y+startY, char, nil, tcell.StyleDefault)
			x += runewidth.RuneWidth(char)
		}
		x = 0
		if y+startY >= boundY {
			continue
		}
		y += 1
	}
	return nil
}
