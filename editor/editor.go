package editor

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	Tabs  []Tab
	TabId int

	Keymap Keymap

	Screen tcell.Screen
}

func (e *Editor) tab() *Tab {
	return &e.Tabs[e.TabId]
}

func (e *Editor) Setup() {
	e.Tabs = []Tab{{
		Buffer:   &buffer.SliceBuffer{},
		FilePath: "",
		Mode:     ModeNormal,
	}} // TEMP:
	e.tab().Buffer.Write([][]rune{
		[]rune("hello world"),
		[]rune("hello tubre"),
	}) // TEMP:

	e.Keymap = defaultKeymap
	// TODO: Load custom config and Keymap

	var err error

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
