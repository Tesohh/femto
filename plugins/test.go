package plugins

import (
	"github.com/Tesohh/femto/editor"
	"github.com/gdamore/tcell/v2"
)

type MyPlugin struct{}

func (p *MyPlugin) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{}
}

func (p *MyPlugin) Startup(e *editor.Editor) error {
	return nil
}

func (p *MyPlugin) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	return nil
}

func (p *MyPlugin) Draw(e *editor.Editor) error {
	return nil
}
