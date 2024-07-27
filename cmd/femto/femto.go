package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Tesohh/femto/editor"
)

func logErr(err error) {
	if err, ok := err.(editor.FemtoError); ok {
		// TODO: do different things based on error level
		slog.Log(context.TODO(), err.LogLevel, err.Error())
	} else {
		slog.Error(err.Error())
	}
}

func main() {
	f, err := os.OpenFile("femto.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(f, nil))
	slog.SetDefault(logger)

	e := editor.Editor{}
	e.Setup()

	for {
		err := e.Update()
		if err != nil {
			logErr(err)
			continue
		}
		err = e.Draw()
		if err != nil {
			logErr(err)
			continue
		}
	}

}
