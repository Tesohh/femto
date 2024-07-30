package editor

import "os"

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
}
