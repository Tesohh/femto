package plugins

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Tesohh/femto/buffer"
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
			Func: func(e *editor.Editor, args ...string) error {
				win := e.Win()

				var path string
				if len(args) == 0 || args[0] == "" {
					if win.FilePath == "" {
						return ErrWindowIsScratchpad
					}
					path = win.FilePath
				} else {
					path = args[0]
				}

				dir := filepath.Dir(path)
				err := os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return ErrFailedToCreateDirs.Context(err.Error())
				}

				b, err := e.Buf().Read1D()
				if err != nil {
					return err
				}

				err = os.WriteFile(path, []byte(string(b)), 0644)
				if err != nil {
					return err
				}

				e.Screen.PostEvent(&editor.CommandBarEvent{
					Msg:  fmt.Sprintf("wrote file %s", path),
					Time: time.Now(),
				})

				return nil
			},
		},
		"w": editor.Alias("write"),
		"edit": {
			Name: "Opens specified file for editing in current file",
			Func: func(e *editor.Editor, args ...string) error {
				if len(args) == 0 || args[0] == "" {
					return editor.ErrArgumentMissing
				}

				path := args[0]

				b, err := os.ReadFile(path)
				if err != nil {
					// no big deal
					slog.Warn(err.Error())
					b = []byte{}
					e.Screen.PostEvent(&editor.CommandBarEvent{
						Msg:  fmt.Sprintf("file %s not found, will be created when writing", path),
						Time: time.Now(),
					})
				}

				lines := strings.Split(string(b), "\n")

				rs := [][]rune{}
				for _, line := range lines {
					rs = append(rs, []rune(line))
				}

				buf := buffer.SliceBuffer{}
				err = buf.Write(rs)
				if err != nil {
					return err
				}

				win := e.Tab().GetWindow("main")
				win.FilePath = path
				win.Buffer = &buf

				return nil
			},
		},
		"e": editor.Alias("edit"),
	},
}
