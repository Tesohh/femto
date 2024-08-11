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
			":": "command.focus",
		},
	})

	e.RegisterCommandMap(map[string]editor.Command{
		"command.focus": {
			Func: func(e *editor.Editor) error {
				e.FocusWindow("commandbar")
				p.w.Buffer.Write([][]rune{{':'}})
				e.Win().Mode = "insert" // no idea why p.w.Mode doesn't do the trick
				p.w.Buffer.ForceRight(1)
				return nil
			},
		},
	})

	p.w = e.RegisterWindow(editor.Window{
		Id:        "commandbar",
		Alignment: editor.AlignmentBottom,
		Size:      1,
		Priority:  3,
		Shown:     true,
		Flags:     editor.WindowFlagUnfocusable,

		Keymap: humankey.HumanKeymap{
			"insert": {
				"esc":   "command.exit",
				"enter": "command.execute",
			},
		},
		Commands: map[string]editor.Command{
			"command.exit": {
				Func: func(e *editor.Editor) error {
					e.Tab().FocusWindow(e, "editor")
					p.w.Buffer.Write([][]rune{{}})
					return nil
				},
			},
			"command.execute": {
				Func: func(e *editor.Editor) error {
					runes, cerr := p.w.Buffer.Read()
					if cerr != nil {
						return cerr
					}

					id := string(runes[0])
					id = strings.Trim(id, ": ")

					e.RunCommand("command.exit")

					cerr = e.RunCommand(id)
					if cerr != nil {
						e.Screen.PostEvent(&editor.EventCaught{}) // send a empty event to force a screen refresh
						return cerr
					}
					return nil
				},
			},
		},
	})

	p.w.Buffer.Write([][]rune{[]rune("Welcome to femto")})

	return err
}

func (p *CommandMode) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	return nil
}

func (p *CommandMode) Draw(e *editor.Editor) error {
	return nil
}
