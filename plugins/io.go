package plugins

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/Tesohh/femto/editor"
)

var (
	ErrWindowIsScratchpad = editor.FemtoError{Message: "cannot write, window is a scratchpad", LogLevel: slog.LevelWarn}
	ErrWindowIsReadonly   = editor.FemtoError{Message: "cannot write, window is readonly", LogLevel: slog.LevelWarn}
	ErrFailedToCreateDirs = editor.FemtoError{Message: "failed to create directories", LogLevel: slog.LevelError}
)

var Io = editor.DumbPlugin{
	Info: editor.PluginInfo{
		Id:          "femto.io",
		Author:      "femto",
		Name:        "Input/Output",
		Description: "Adds basic commands for IO such as :w[rite]",
	},
	Commands: map[string]editor.Command{
		"write": {
			Name: "Write file",
			Func: func(e *editor.Editor) error {
				win := e.Win()
				if win.FilePath == "" {
					return ErrWindowIsScratchpad
				}

				dir := filepath.Dir(win.FilePath)
				err := os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return ErrFailedToCreateDirs.Context(err.Error())
				}

				b, err := e.Buf().Read1D()
				if err != nil {
					return err
				}

				err = os.WriteFile(win.FilePath, []byte(string(b)), 0644)
				if err != nil {
					return err
				}

				e.Screen.PostEvent(&editor.CommandBarEvent{
					Msg:  fmt.Sprintf("wrote file %s", win.FilePath),
					Time: time.Now(),
				})

				return nil
			},
		},
		"w": editor.Alias("write"),
	},
}
