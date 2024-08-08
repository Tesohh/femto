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

	Themes map[string]Theme
	Theme  Theme // for simplicity's (and performance) sake, since we won't change theme often, we don't save the id but just the theme itself here

	Screen             tcell.Screen
	Windows            []Window
	FocusedWindowIndex int // if set to something that isn't -1, overrides the Tab's CurrentWindowId
}

func (e *Editor) Tab() *Tab {
	return &e.Tabs[e.TabId]
}
func (e *Editor) Win() *Window {
	if e.FocusedWindowIndex < 0 {
		return &e.Tabs[e.TabId].Windows[e.Tab().FocusedWindowIndex]
	} else {
		return &e.Windows[e.FocusedWindowIndex]
	}
}
func (e *Editor) Buf() buffer.Buffer {
	return e.Win().Buffer
}

func (e *Editor) Setup() {
	e.Tabs = []Tab{{}} // TEMP:

	e.Tabs[0].RegisterWindow(Window{
		Alignment: AlignmentCenter,
		Shown:     true,
	}) // TEMP:

	e.FocusedWindowIndex = -1

	e.Buf().Write([][]rune{
		[]rune("hello world"),
		[]rune("hello tubre and zernez"),
		[]rune("oh slicebufer"),
		[]rune("        java"),
	}) // TEMP:

	e.Commands = Commands
	e.Themes = map[string]Theme{}

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
