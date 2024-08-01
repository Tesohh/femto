package editor

import (
	"os"
	"unicode"

	"github.com/Tesohh/femto/femath"
)

type CommandFunc func(e *Editor) error

func (e *Editor) RunCommand(id string) error {
	cmd, ok := Commands[id]
	if !ok {
		return ErrNoCommandFound.Context(id)
	}

	return cmd.Func(e)
}

type Command struct {
	Name        string
	Description string // if empty, takes the Name as Description
	Func        CommandFunc
}

var Commands = map[string]Command{
	"write": {
		Name: "Write file",
		Func: func(e *Editor) error {
			panic("not implemented")
		},
	},
	"quit": {
		Name: "quit editor",
		Func: func(e *Editor) error {
			e.Screen.Fini()
			os.Exit(0)
			return nil
		},
	},
	"left": {
		Func: func(e *Editor) error {
			e.tab().Buffer.Left(1)
			return nil
		},
	},
	"down": {
		Func: func(e *Editor) error {
			e.tab().Buffer.Down(1)
			return nil
		},
	},
	"up": {
		Func: func(e *Editor) error {
			e.tab().Buffer.Up(1)
			return nil
		},
	},
	"right": {
		Func: func(e *Editor) error {
			e.tab().Buffer.Right(1)
			return nil
		},
	},
	"bigLeft": {
		Func: func(e *Editor) error {
			line := e.tab().Buffer.Line()
			x := 0
			for ; x < len(line); x++ {
				if !unicode.IsSpace(line[x]) {
					break
				}
			}

			e.tab().Buffer.GoTo(femath.Vec2{X: x, Y: e.tab().Buffer.Pos().Y})
			return nil
		},
	},
	"bigRight": {
		Func: func(e *Editor) error {
			e.tab().Buffer.GoTo(femath.Vec2{X: len(e.tab().Buffer.Line()), Y: e.tab().Buffer.Pos().Y})
			return nil
		},
	},
}
