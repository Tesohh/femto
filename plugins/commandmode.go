package plugins

import (
	"strings"
	"time"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type CommandBar struct {
	w *editor.Window
}

func (p *CommandBar) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:     "femto.commandbar",
		Author: "femto",
		Name:   "Command bar",
	}
}

func (p *CommandBar) Startup(e *editor.Editor) error {
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
				e.Win().StyleSections = []editor.StyleSection{}
				e.Win().Mode = "insert" // no idea why p.w.Mode doesn't do the trick
				p.w.Buffer.ForceRight(1)
				return nil
			},
		},
		"command.test": {
			Func: func(e *editor.Editor) error {
				e.Screen.PostEvent(&editor.CommandBarEvent{
					Msg:   "Cissy",
					Style: e.Theme.Error,
					Time:  time.Now(),
				})
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

func (p *CommandBar) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	if event, ok := event.(*editor.CommandBarEvent); ok {
		p.w.Buffer.Write([][]rune{[]rune(event.Msg)})
		// p.w.StyleSections = []editor.StyleSection{}
		// p.w.StyleSections = append(p.w.StyleSections, editor.StyleSection{
		// 	Y:      0,
		// 	StartX: 0,
		// 	EndX:   len(event.Msg),
		// 	Style:  event.Style,
		// })
		return &editor.EventCaught{}
	}
	return nil
}

func (p *CommandBar) Draw(e *editor.Editor) error {
	return nil
}
