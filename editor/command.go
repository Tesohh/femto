package editor

import (
	"os"
)

type CommandFunc func(e *Editor) error

func (e *Editor) RunCommand(id string) error {
	editorCmd, editorOk := e.Commands[id]
	windowCmd, windowOk := e.Win().Commands[id]

	if !editorOk && !windowOk {
		return ErrNoCommandFound.Context(id)
	}

	if windowOk {
		return windowCmd.Func(e)
	} else {
		return editorCmd.Func(e)
	}
}

type Command struct {
	Name        string
	Description string // if empty, takes the Name as Description
	Public      bool   // if public, can be executed in command mode
	Func        CommandFunc
}

func Alias(cmd string) Command {
	return Command{
		Func: func(e *Editor) error {
			return e.RunCommand(cmd)
		}}
}

var Commands = map[string]Command{
	"noop": {
		Name: "no operation",
		Func: func(e *Editor) error {
			return nil
		},
	},
	"normal": {
		Name: "Normal mode",
		Func: func(e *Editor) error {
			e.Win().Mode = "normal"
			return nil
		},
	},
	"write": {
		Name: "Write file",
		Func: func(e *Editor) error {
			panic("not implemented")
		},
	},
	"w": Alias("write"),
	"quit": {
		Name: "quit editor",
		Func: func(e *Editor) error {
			e.Screen.Fini()
			os.Exit(0)
			return nil
		},
	},
	"q": Alias("quit"),
}
