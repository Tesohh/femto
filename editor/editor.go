package editor

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	Tabs  []Tab
	TabId int

	Keymap   humankey.InternalKeymap
	Commands map[string]Command
	Plugins  []Plugin

	Screen  tcell.Screen
	Windows []Window
}

func (e *Editor) Tab() *Tab {
	return &e.Tabs[e.TabId]
}
func (e *Editor) Buf() buffer.Buffer { // dont need to pointer it: interfaces ARE pointers
	return e.Tabs[e.TabId].Buffer
}

func (e *Editor) Setup() {
	e.Tabs = []Tab{{
		Buffer:   &buffer.SliceBuffer{},
		FilePath: "",
		Mode:     "normal",
	}} // TEMP:
	e.Tab().Buffer.Write([][]rune{
		[]rune("hello world"),
		[]rune("hello tubre and zernez"),
		[]rune("oh slicebufer"),
		[]rune("        java"),
	}) // TEMP:

	e.Windows = []Window{
		{
			Alignment: AlignmentBottom,
			Size:      3,
			Priority:  3,
			Shown:     true,
			Flags:     WindowFlagHasBorder,
			Content:   &buffer.SliceBuffer{},
		},
		{
			Alignment:   AlignmentLeft,
			Size:        15,
			Priority:    1,
			Shown:       true,
			Flags:       WindowFlagHasBorder,
			Content:     &buffer.SliceBuffer{},
			BorderStyle: tcell.StyleDefault.Dim(true),
		},
		{
			Alignment:   AlignmentRight,
			Size:        15,
			Priority:    1,
			Shown:       true,
			Flags:       WindowFlagHasBorder,
			Content:     &buffer.SliceBuffer{},
			BorderStyle: tcell.StyleDefault.Dim(true),
		},
		{
			Alignment: AlignmentTop,
			Size:      3,
			Priority:  0,
			Shown:     true,
			Flags:     WindowFlagHasBorder,
			Content:   &buffer.SliceBuffer{},
		},
	}
	e.Windows[0].Content.Write([][]rune{
		[]rune("cissy"),
		[]rune("la bottombar"),
	})
	e.Windows[1].Content.Write([][]rune{
		[]rune("sinistro"),
		[]rune("sinistro"),
	})
	e.Windows[2].Content.Write([][]rune{
		[]rune("destroooooooooo"),
		[]rune("destro"),
	})
	e.Windows[3].Content.Write([][]rune{
		[]rune("cissy.go"),
		[]rune("func main() >"),
	})

	e.Commands = Commands

	var err error
	e.Keymap, err = defaultKeymap.ToInternal()
	if err != nil {
		panic(err)
	}
	// TODO: Load custom config and Keymap

	e.Screen, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = e.Screen.Init()
	if err != nil {
		panic(err)
	}
	e.Screen.Clear()
}
