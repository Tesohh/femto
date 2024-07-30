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

	Screen tcell.Screen
}

func (e *Editor) tab() *Tab {
	return &e.Tabs[e.TabId]
}

func (e *Editor) Setup() {
	e.Tabs = []Tab{{
		Buffer:   &buffer.SliceBuffer{},
		FilePath: "",
		Mode:     "normal",
	}} // TEMP:
	e.tab().Buffer.Write([][]rune{
		[]rune("hello world"),
		[]rune("hello tubre"),
	}) // TEMP:

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
