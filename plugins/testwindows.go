package plugins

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/Tesohh/femto/editor"
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
	e.Windows = []editor.Window{
		{
			Alignment: editor.AlignmentBottom,
			Size:      3,
			Priority:  3,
			Shown:     true,
			Flags:     editor.WindowFlagHasBorder,
			Buffer:    &buffer.SliceBuffer{},
		},
		{
			Alignment:   editor.AlignmentLeft,
			Size:        15,
			Priority:    1,
			Shown:       true,
			Flags:       editor.WindowFlagHasBorder,
			Buffer:      &buffer.SliceBuffer{},
			BorderStyle: tcell.StyleDefault.Dim(true),
		},
		{
			Alignment:   editor.AlignmentRight,
			Size:        15,
			Priority:    1,
			Shown:       true,
			Flags:       editor.WindowFlagHasBorder,
			Buffer:      &buffer.SliceBuffer{},
			BorderStyle: tcell.StyleDefault.Dim(true),
		},
		{
			Alignment: editor.AlignmentTop,
			Size:      3,
			Priority:  0,
			Shown:     true,
			Flags:     editor.WindowFlagHasBorder,
			Buffer:    &buffer.SliceBuffer{},
		},
		{
			Alignment: editor.AlignmentCenter,
			Priority:  0,
			Shown:     true,
			Buffer:    &buffer.SliceBuffer{},
		},
	}
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
		[]rune("func main() >"),
	})
	e.Windows[4].Buffer.Write([][]rune{
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
		[]rune("ciojweofijwefiwefoiwejfoiwej"),
	})
	return nil
}
func (p *TestWindowsPlugin) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	return nil
}
func (p *TestWindowsPlugin) Draw(e *editor.Editor) error {
	return nil
}
