package main

import (
	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/plugins"
)

var pluginsList = []editor.Plugin{
	&plugins.Movement,
	&plugins.InsertMode{},
	&plugins.TestWindowsPlugin{},
}
