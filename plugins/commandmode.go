package plugins

import (
	"log/slog"
	"strings"
	"time"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type CommandBar struct {
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
				e.Win().Buffer.Write([][]rune{{':'}})
				e.Win().StyleSections = []editor.StyleSection{}
				e.Win().Mode = "insert"
				e.Win().Buffer.ForceRight(1)
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

	e.RegisterWindow(editor.Window{
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
					e.Tab().FocusWindow(e, "main")
					e.GetWindow("commandbar").Buffer.Write([][]rune{{}})
					return nil
				},
			},
			"command.execute": {
				Func: func(e *editor.Editor) error {
					runes, cerr := e.GetWindow("commandbar").Buffer.Read()
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

	e.GetWindow("commandbar").Buffer.Write([][]rune{[]rune("Welcome to femto")})

	return err
}

func (p *CommandBar) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	if event, ok := event.(*editor.CommandBarEvent); ok {
		w := e.GetWindow("commandbar")
		w.Buffer.Write([][]rune{[]rune(event.Msg)})
		w.StyleSections = []editor.StyleSection{}
		w.StyleSections = append(w.StyleSections, editor.StyleSection{
			Y:      0,
			StartX: 0,
			EndX:   len(event.Msg),
			Style:  event.Style,
		})
		return &editor.EventCaught{}
	}
	return nil
}

func (p *CommandBar) Draw(e *editor.Editor) error {
	return nil
}

func CommandBarTryPushMessage(e *editor.Editor, err error) {
	w := e.GetWindow("commandbar")
	if w == nil {
		slog.Warn("commandbar window not found")
		return
	}

	var style tcell.Style
	if err, ok := err.(editor.FemtoError); ok {

		switch err.LogLevel {
		case slog.LevelInfo:
			return
		case slog.LevelWarn:
			style = e.Theme.Warn
		case slog.LevelError:
			style = e.Theme.Error
		}

	}

	e.Screen.PostEvent(&editor.CommandBarEvent{
		Msg:   err.Error(),
		Style: style,
		Time:  time.Now(),
	})
}
