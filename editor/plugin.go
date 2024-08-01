package editor

import (
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

func (e *Editor) RegisterCommandMap(cmds map[string]Command) {
	// names := strings.Split(cmd.Name, " ")
	// for i, s := range names[1:] {
	// 	names[i] = strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	// }
	//
	// name := strings.Join(names, "")
	// id := fmt.Sprintf("%s.%s", namespace, name)

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
	Startup(*Editor) error
	Draw(*Editor) error
	Update(*Editor, tcell.Event) error
}
