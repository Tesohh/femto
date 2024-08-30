package plugins

import (
	"fmt"
	"log/slog"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/femath"
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

// var VisualMode = editor.DumbPlugin{
// 	Info: editor.PluginInfo{
// 		Id:          "femto.visualmode",
// 		Author:      "femto",
// 		Name:        "Visual Mode",
// 		Description: "Adds the visual mode to control your selection",
// 	},
// 	Commands: map[string]editor.Command{
// 		"visual.mode": {
// 			Name: "enter visual mode",
// 			Func: func(e *editor.Editor, _ ...string) error {
// 				win := e.Win()
// 				win.Mode = "visual"
// 				win.Selection.Start = win.Buffer.Pos()
// 				win.Selection.End = win.Buffer.Pos()
// 				return nil
// 			},
// 		},
// 	},
// 	Keymap: map[string]map[string]string{},
// 	Themes: map[string]editor.Theme{},
// }

type VisualMode struct{}

func (p *VisualMode) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:          "femto.visualmode",
		Author:      "femto",
		Name:        "Visual Mode",
		Description: "Adds the visual mode to control your selection",
	}
}

func (p *VisualMode) Startup(e *editor.Editor) error {
	e.RegisterCommandMap(map[string]editor.Command{
		"visual.mode": {
			Name: "enter visual mode",
			Func: func(e *editor.Editor, _ ...string) error {
				win := e.Win()
				win.Mode = "visual"
				win.Selection.Start = win.Buffer.Pos()
				win.Selection.End = win.Buffer.Pos()
				return nil
			},
		},
		"visual.linemode": {
			Name: "enter visual linewise mode",
			Func: func(e *editor.Editor, _ ...string) error {
				win := e.Win()
				win.Mode = "viline"
				win.Selection.Start = femath.Vec2{X: 0, Y: win.Buffer.Pos().Y}
				win.Selection.End = femath.Vec2{X: 9999, Y: win.Buffer.Pos().Y} // FIX: if you have more than 9999 chars...
				return nil
			},
		},
		"visual.tonormal": {
			Func: func(e *editor.Editor, _ ...string) error {
				win := e.Win()
				win.Mode = "normal"
				win.Selection.Start = femath.Vec2{X: 0, Y: 0}
				win.Selection.End = femath.Vec2{X: 0, Y: 0}
				return nil
			},
		},
	})

	e.RegisterKeymap(humankey.HumanKeymap{
		"normal": {
			"v": "visual.mode",
			"V": "visual.linemode",
		},
		"visual": {
			"h":     "left",
			"left":  "left",
			"j":     "down",
			"down":  "down",
			"k":     "up",
			"up":    "up",
			"l":     "right",
			"right": "right",
			"H":     "bigLeft",
			"L":     "bigRight",
			"esc":   "visual.tonormal",
		},
		"viline": {
			"j":    "down",
			"down": "down",
			"k":    "up",
			"up":   "up",
			"esc":  "visual.tonormal",
		},
	})

	return nil
}

func (p *VisualMode) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	win := e.Win()
	if win.Mode == "visual" {
		win.Selection.End = win.Buffer.Pos()
	} else if win.Mode == "viline" {
		win.Selection.End.Y = win.Buffer.Pos().Y
	}

	slog.Info(fmt.Sprintf("%#v", win.Selection))
	return nil
}

func (p *VisualMode) Draw(e *editor.Editor) error {
	return nil
}
