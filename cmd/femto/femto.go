package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/plugins"
)

func main() {
	f, err := os.OpenFile("femto.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(f, nil))
	slog.SetDefault(logger)

	e := editor.Editor{}
	e.Setup()
	e.Plugins = pluginsList // TODO: Load only enabled plugins + mandatory plugins

	for _, p := range e.Plugins {
		err := p.Startup(&e)
		if err != nil {
			panic(err)
		}
	}

	e.Theme = e.Themes["rosepine.main"]

	for {
		err := e.Update()
		p := recover()
		if p != nil {
			slog.Info(fmt.Sprintf("%v", p))
		}
		if err != nil {
			logErr(err)
			plugins.CommandBarTryPushMessage(&e, err)
			continue
		}

		err = e.Draw()
		if err != nil {
			logErr(err)
			plugins.CommandBarTryPushMessage(&e, err)
			continue
		}
	}

}

func logErr(err error) {
	if ferr, ok := err.(editor.FemtoError); ok {
		slog.Log(context.TODO(), ferr.LogLevel, ferr.Error())
	} else {
		slog.Error(err.Error())
	}
}
