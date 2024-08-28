package editor

import (
	"fmt"
	"log/slog"
	"os"
)

type CommandFunc func(e *Editor, args ...string) error

func (e *Editor) RunCommand(id string, args ...string) error {
	editorCmd, editorOk := e.Commands[id]
	windowCmd, windowOk := e.Win().Commands[id]

	if !editorOk && !windowOk {
		return ErrNoCommandFound.Context(id)
	}

	slog.Info(fmt.Sprintf("running command %s with args %v", id, args))

	if windowOk {
		return windowCmd.Func(e, args...)
	} else {
		return editorCmd.Func(e, args...)
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
		Func: func(e *Editor, args ...string) error {
			return e.RunCommand(cmd, args...)
		}}
}

var Commands = map[string]Command{
	"noop": {
		Name: "no operation",
		Func: func(e *Editor, args ...string) error {
			return nil
		},
	},
	"normal": {
		Name: "Normal mode",
		Func: func(e *Editor, args ...string) error {
			e.Win().Mode = "normal"
			return nil
		},
	},
	"quit": {
		Name: "quit editor",
		Func: func(e *Editor, args ...string) error {
			e.Screen.Fini()
			os.Exit(0)
			return nil
		},
	},
	"q": Alias("quit"),
}
