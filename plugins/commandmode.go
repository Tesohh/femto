package plugins

import (
	"strings"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

// FIX: tear this code down, and make it an interactive window with keymap/command overrides.
// just make it so that, in insert mode for this window
// esc closes the window and goes to normal
// and enter executes instead

// TEMP: ahead!!

type CommandMode struct {
	w *editor.Window
}

func (p *CommandMode) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:     "femto.commandmode",
		Author: "femto",
		Name:   "Command Mode",
	}
}

func (p *CommandMode) Startup(e *editor.Editor) error {
	err := e.RegisterKeymap(humankey.HumanKeymap{
		"normal": {
			":": "commandmode",
		},
		"command": {
			"esc":   "command.tonormal",
			"enter": "command.execute",
		},
	})

	e.RegisterCommandMap(map[string]editor.Command{
		"commandmode": {
			Func: func(e *editor.Editor) error {
				p.w.Shown = true
				e.Win().Mode = "command"
				return p.w.Buffer.Write([][]rune{{}})
			},
		},
		"command.tonormal": {
			Func: func(e *editor.Editor) error {
				p.w.Shown = false
				e.Win().Mode = "normal"
				return p.w.Buffer.Write([][]rune{{}}) // clear
			},
		},
		"command.execute": {
			Func: func(e *editor.Editor) error {
				r, err := p.w.Buffer.Read()
				if err != nil {
					return err
				}

				s := string(r[0])
				s = strings.Trim(s, ": ")

				err = e.RunCommand(s)
				if err != nil {
					return err
				}
				return e.RunCommand("command.tonormal")
			},
		},
	})

	p.w = e.RegisterWindow(editor.Window{
		Id:        "commandbar",
		Alignment: editor.AlignmentBottom,
		Size:      1,
		Priority:  3,
		Flags:     editor.WindowFlagUnfocusable,
	})

	return err
}

func (p *CommandMode) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	if e.Win().Mode != "command" {
		return nil
	}

	if event, ok := event.(*tcell.EventKey); ok {
		if event.Key() == tcell.KeyRune {
			p.w.Buffer.Insert(p.w.Buffer.Pos(), event.Rune())
			p.w.Buffer.ForceRight(1)
			return &editor.EventCaught{}
		}
	}

	return nil
}

func (p *CommandMode) Draw(e *editor.Editor) error {
	if e.Win().Mode != "command" {
		return nil
	}

	return nil
}
