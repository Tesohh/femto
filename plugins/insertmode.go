package plugins

import (
	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type InsertMode struct{}

var info = editor.PluginInfo{
	Id:     "femto.insertmode",
	Author: "femto",
	Name:   "Insert mode",
}

var keymap = humankey.HumanKeymap{
	"normal": {
		"i": "insert.i",
		"a": "insert.a",
	},
	"insert": {
		"esc":        "insert.tonormal",
		"backspace":  "insert.backspace",
		"backspace2": "insert.backspace",
	},
}

var commands = map[string]editor.Command{
	"insert.i": {
		Func: func(e *editor.Editor, args ...string) error {
			e.Win().Mode = "insert"
			return nil
		},
	},
	"insert.a": {
		Func: func(e *editor.Editor, args ...string) error {
			e.Buf().ForceRight(1)
			e.Win().Mode = "insert"
			return nil
		},
	},
	"insert.tonormal": {
		Func: func(e *editor.Editor, args ...string) error {
			e.Win().Mode = "normal"
			e.Buf().Left(1)
			return nil
		},
	},
	"insert.backspace": {
		Func: func(e *editor.Editor, args ...string) error {
			if e.Buf().Pos().X == 0 {
				return nil
			}
			e.Buf().Left(1)
			e.Buf().Delete(e.Buf().Pos())
			return nil
		},
	},
}

func (p *InsertMode) GetInfo() editor.PluginInfo {
	return info
}

func (p *InsertMode) Startup(e *editor.Editor) error {
	e.RegisterCommandMap(commands)
	return e.RegisterKeymap(keymap)
}

func (p *InsertMode) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	if e.Win().Mode != "insert" {
		return nil
	}

	if event, ok := event.(*tcell.EventKey); ok {
		if event.Key() == tcell.KeyRune {
			e.Buf().Insert(e.Buf().Pos(), event.Rune())
			e.Buf().ForceRight(1)
			return &editor.EventCaught{}
		}
	}

	return nil
}

func (p *InsertMode) Draw(e *editor.Editor) error {
	return nil
}
