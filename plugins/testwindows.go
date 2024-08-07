package plugins

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type TestWindowsPlugin struct{}

func (p *TestWindowsPlugin) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:     "femto.temp.windowstest",
		Author: "femto",
		Name:   "temporary windows test",
	}
}
func (p *TestWindowsPlugin) Startup(e *editor.Editor) error {
	e.RegisterKeymap(humankey.HumanKeymap{
		"normal": {
			"1": "test_moving_window",
		},
	})
	e.RegisterCommandMap(map[string]editor.Command{
		"test_moving_window": {
			Func: func(e *editor.Editor) error {
				e.FocusedWindowIndex = 1
				return nil
			},
		},
	})
	e.RegisterWindow(editor.Window{
		Alignment: editor.AlignmentBottom,
		Size:      1,
		Priority:  3,
		Shown:     true,
	})
	e.RegisterWindow(editor.Window{
		Alignment:   editor.AlignmentLeft,
		Size:        15,
		Priority:    1,
		Shown:       true,
		Flags:       editor.WindowFlagHasBorder,
		BorderStyle: tcell.StyleDefault.Dim(true),
	})
	e.RegisterWindow(editor.Window{
		Alignment:   editor.AlignmentRight,
		Size:        15,
		Priority:    1,
		Shown:       true,
		Flags:       editor.WindowFlagHasBorder,
		BorderStyle: tcell.StyleDefault.Dim(true),
	})
	e.RegisterWindow(editor.Window{
		Alignment: editor.AlignmentTop,
		Size:      2,
		Priority:  0,
		Shown:     true,
		Buffer:    &buffer.SliceBuffer{},
	})

	e.Windows[0].Buffer.Write([][]rune{
		[]rune("cissy"),
		[]rune("la bottombar"),
	})
	e.Windows[1].Buffer.Write([][]rune{
		[]rune("sinistro"),
		[]rune("sinistro"),
	})
	e.Windows[2].Buffer.Write([][]rune{
		[]rune("destroooooooooo"),
		[]rune("destro"),
	})
	e.Windows[3].Buffer.Write([][]rune{
		[]rune("cissy.go"),
	})
	return nil
}
func (p *TestWindowsPlugin) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	return nil
}
func (p *TestWindowsPlugin) Draw(e *editor.Editor) error {
	return nil
}
