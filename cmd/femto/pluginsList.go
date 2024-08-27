package main

import (
	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/plugins"
)

var pluginsList = []editor.Plugin{
	&plugins.Movement,
	&plugins.Io,
	&plugins.RosePine,
	&plugins.InsertMode{},
	&plugins.CommandBar{},
	&plugins.StatusBar{},
	// &plugins.TestWindowsPlugin{},
}
