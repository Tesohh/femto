package plugins

import (
	"fmt"

	"github.com/Tesohh/femto/editor"
	"github.com/gdamore/tcell/v2"
)

type StatusBar struct {
	w *editor.Window
}

func (s *StatusBar) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:     "femto.statusbar",
		Author: "femto",
		Name:   "Status bar",
	}
}

func (s *StatusBar) Startup(e *editor.Editor) error {
	s.w = e.RegisterWindow(editor.Window{
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
	style := tcell.StyleDefault
	switch e.Win().Mode {
	case "normal":
		mode = "NORMAL"
		style = style.Background(tcell.Color(0x303030)) // TEMP: should read from theme
	case "insert":
		mode = "INSERT"
		style = style.Background(tcell.Color(0x303030)) // TEMP: should read from theme
	}

	_ = style // TEMP:

	str := fmt.Sprintf(" %s ", mode)

	s.w.Buffer.Write([][]rune{[]rune(str)})
	return nil
}
