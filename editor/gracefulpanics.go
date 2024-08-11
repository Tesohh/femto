package editor

import (
	"fmt"
	"log/slog"
)

// Panicking without finalizing the screen causes really weird behaviour,
// So defer this at the top of every main loop function (not for plugins)
func GracefulPanic(e *Editor) {
	if r := recover(); r != nil {
		e.Screen.Fini()
		slog.Error(fmt.Sprintf("PANIC PANIC PANIC %#v", r))
		panic(r)
	}
}
