package plugins

import (
	"fmt"

	"github.com/Tesohh/femto/editor"
	"github.com/gdamore/tcell/v2"
)

type StatusBar struct {
}

func (s *StatusBar) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:     "femto.statusbar",
		Author: "femto",
		Name:   "Status bar",
	}
}

func (s *StatusBar) Startup(e *editor.Editor) error {
	e.RegisterWindow(editor.Window{
		Id:        "statusbar",
		Alignment: editor.AlignmentBottom,
		Priority:  3,
		Shown:     true,
		Size:      1,
		Flags:     editor.WindowFlagReadonly | editor.WindowFlagUnfocusable,
	})
	return nil
}

func (s *StatusBar) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	return nil
}

func (s *StatusBar) Draw(e *editor.Editor) error {
	mode := ""
	style := tcell.StyleDefault.Bold(true)
	switch e.Win().Mode {
	case "normal":
		mode = "NORMAL"
		style = style.Background(e.Theme.NormalModeAccent)
	case "insert":
		mode = "INSERT"
		style = style.Background(e.Theme.InsertModeAccent)
	case "visual":
		mode = "VISUAL"
		style = style.Background(e.Theme.VisualModeAccent)
	case "viline":
		mode = "VILINE"
		style = style.Background(e.Theme.VisualModeAccent)
	}

	str := fmt.Sprintf(" %s ", mode)

	w := e.GetWindow("statusbar")
	w.StyleSections = []editor.StyleSection{}
	w.StyleSections = append(w.StyleSections, editor.StyleSection{
		Y:      0,
		StartX: 0,
		EndX:   len(" NORMAL "),
		Style:  style,
	})
	w.Buffer.Write([][]rune{[]rune(str)})
	return nil
}
