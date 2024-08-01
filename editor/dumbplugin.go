package editor

import (
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

// A plugin that only has a Startup function,
// and can contribute Commands and Keymap
//
// ideally only use this, unless you REALLY need to access the main loop
type DumbPlugin struct {
	Info     PluginInfo
	Commands map[string]Command
	Keymap   humankey.HumanKeymap
}

func (p *DumbPlugin) GetInfo() PluginInfo {
	return p.Info
}

func (p *DumbPlugin) Startup(e *Editor) error {
	if p.Commands != nil {
		e.RegisterCommandMap(p.Commands)
	}

	if p.Keymap != nil {
		e.RegisterKeymap(p.Keymap)
	}

	return nil
}

func (p *DumbPlugin) Draw(e *Editor) error {
	return nil
}

func (p *DumbPlugin) Update(e *Editor, event tcell.Event) error {
	return nil
}
