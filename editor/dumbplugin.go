package editor

import (
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

// A plugin that only has a Startup function,
// and can contribute Commands and Keymap
//
// ideally only use **this**, unless you REALLY need to access the main loop
type DumbPlugin struct {
	Info     PluginInfo
	Commands map[string]Command // if it's a third party plugin, please prefix Commands with your plugin id
	Keymap   humankey.HumanKeymap
	Themes   map[string]Theme
}

func (p *DumbPlugin) GetInfo() PluginInfo {
	return p.Info
}

func (p *DumbPlugin) Startup(e *Editor) error {
	if p.Commands != nil {
		e.RegisterCommandMap(p.Commands)
	}

	var err error
	if p.Keymap != nil {
		err = e.RegisterKeymap(p.Keymap)
	}

	if p.Themes != nil {
		e.RegisterThemeMap(p.Themes)
	}

	return err
}

func (p *DumbPlugin) Draw(e *Editor) error {
	return nil
}

func (p *DumbPlugin) Update(e *Editor, event tcell.Event) tcell.Event {
	return nil
}
