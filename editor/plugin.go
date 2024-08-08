package editor

import (
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

func (e *Editor) RegisterThemeMap(themes map[string]Theme) {
	for k, v := range themes {
		e.Themes[k] = v
	}
}

func (e *Editor) RegisterCommandMap(cmds map[string]Command) {
	for k, v := range cmds {
		e.Commands[k] = v
	}
}

func (e *Editor) RegisterKeymap(keymap humankey.HumanKeymap) error {
	internal, err := keymap.ToInternal()
	if err != nil {
		return err
	}

	for mode, keyset := range internal {
		e.Keymap[mode] = append(e.Keymap[mode], keyset...)
	}

	return nil
}

type PluginInfo struct {
	Id          string
	Author      string
	Name        string
	Description string
}

type Plugin interface {
	GetInfo() PluginInfo
	Startup(e *Editor) error
	Update(e *Editor, event tcell.Event) tcell.Event // Plugins can hijack the event by returning a new one. Only do this for errors or for catching the event
	Draw(e *Editor) error
}
