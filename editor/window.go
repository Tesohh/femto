package editor

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func setupWindow(w *Window) *Window {
	w.Mode = "normal"
	w.Sequence = []humankey.InternalKey{}
	if w.Buffer == nil {
		w.Buffer = &buffer.SliceBuffer{}
	}

	return w
}

func (e *Editor) RegisterWindow(w Window) {
	setupWindow(&w)
	e.Windows = append(e.Windows, w)
}

type Alignment uint8

const (
	AlignmentLeft   Alignment = 0
	AlignmentRight  Alignment = 1
	AlignmentTop    Alignment = 2
	AlignmentBottom Alignment = 3
	AlignmentCenter Alignment = 4
)

type WindowFlags uint8

const (
	WindowFlagReadonly    WindowFlags = 1
	WindowFlagInteractive WindowFlags = 2
	WindowFlagHasBorder   WindowFlags = 4
	WindowFlagUnfocusable WindowFlags = 8
)

type Window struct {
	Id string

	Alignment Alignment
	Size      int
	Priority  int

	Shown   bool
	Focused bool
	Flags   WindowFlags

	// buffer stuff
	Buffer   buffer.Buffer // to implement interactivity, you just need to make a type InteractiveBuffer and runtime check if its that typ
	FilePath string        // if left empty, will treat buffer as scratchpad
	Mode     string
	Sequence []humankey.InternalKey

	// TODO: ColorSections
	BorderStyle tcell.Style
}

func (w *Window) Draw(e *Editor, startX int, startY int, boundX int, boundY int, isFocused bool) error {
	text, err := w.Buffer.Read()
	if err != nil {
		return err
	}

	x := 0
	y := 0

	for _, line := range text {
		if w.Alignment == AlignmentRight { // bandaid fix... if it works it works
			x += 2
		}

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

	if isFocused {
		e.Screen.ShowCursor(w.Buffer.Pos().X+startX, w.Buffer.Pos().Y+startY)
	}

	if w.Flags&WindowFlagHasBorder != 0 {
		switch w.Alignment {
		case AlignmentLeft:
			for y := 0; y < boundY; y++ {
				e.Screen.SetContent(boundX-1, y+startY, '│', nil, w.BorderStyle)
			}
		case AlignmentRight:
			for y := 0; y < boundY; y++ {
				e.Screen.SetContent(startX+1, y+startY, '│', nil, w.BorderStyle)
			}
		case AlignmentTop:
			for x := 0; x < boundX; x++ {
				e.Screen.SetContent(x+startX, boundY-1, '─', nil, w.BorderStyle)
			}
		case AlignmentBottom:
			for x := 0; x < boundX; x++ {
				e.Screen.SetContent(x+startX, startY, '─', nil, w.BorderStyle)
			}
		}
	}
	return nil
}
