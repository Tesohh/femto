package plugins

import (
	"unicode"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/femath"
)

var Movement = editor.DumbPlugin{
	Info: editor.PluginInfo{
		Id:     "femto.movement",
		Author: "femto",
		Name:   "Movement",
	},
	Commands: map[string]editor.Command{
		"left": {
			Func: func(e *editor.Editor) error {
				e.Buf().Left(1)
				return nil
			},
		},
		"down": {
			Func: func(e *editor.Editor) error {
				e.Buf().Down(1)
				return nil
			},
		},
		"up": {
			Func: func(e *editor.Editor) error {
				e.Buf().Up(1)
				return nil
			},
		},
		"right": {
			Func: func(e *editor.Editor) error {
				if e.Win().Mode == "insert" {
					e.Buf().ForceRight(1)
				} else {
					e.Buf().Right(1)
				}
				return nil
			},
		},
		"bigLeft": {
			Func: func(e *editor.Editor) error {
				line := e.Buf().Line()
				x := 0
				for ; x < len(line); x++ {
					if !unicode.IsSpace(line[x]) {
						break
					}
				}

				e.Buf().GoTo(femath.Vec2{X: x, Y: e.Buf().Pos().Y})
				return nil
			},
		},
		"bigRight": {
			Func: func(e *editor.Editor) error {
				e.Buf().GoTo(femath.Vec2{X: len(e.Buf().Line()), Y: e.Buf().Pos().Y})
				return nil
			},
		},
	},
	Keymap: map[string]map[string]string{
		"normal": {
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
		},
		"insert": {
			"left":  "left",
			"down":  "down",
			"up":    "up",
			"right": "right",
		},
	},
}
