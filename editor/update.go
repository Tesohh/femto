package editor

import (
	"github.com/gdamore/tcell/v2"
)

func (e *Editor) Update() error {
	event := e.Screen.PollEvent()

	switch event := event.(type) {
	case *tcell.EventKey:
		// TODO: this can't do anything with regular Rune keypresses. Perhaps, what about using the event.Name() as the keymap key?
		id, ok := e.Keymap[e.tab().Mode][event.Key()]
		if !ok {
			return ErrNoKeyAssociated.Context(event.Name())
		}
		cmd, ok := Commands[id]
		if !ok {
			return ErrNoCommandFound.Context(id)
		}
		err := cmd.Func(e)
		if err != nil {
			return err
		}

	case *tcell.EventError:
		return event
	case *tcell.EventResize:
		w, h := event.Size()
		e.Screen.SetSize(w, h)
		e.Screen.Sync()
	}

	return nil
}
