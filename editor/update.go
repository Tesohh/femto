package editor

import (
	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

func (e *Editor) Update() error {
	event := e.Screen.PollEvent()

	switch event := event.(type) {
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEsc && e.tab().Mode == "normal" {
			// clear sequence with esc in normal mode
			e.tab().Sequence = []humankey.InternalKey{}
			return nil
		}

		key := humankey.InternalKey{
			Key:     event.Key(),
			Rune:    event.Rune(),
			ModMask: event.Modifiers(),
		}

		e.tab().Sequence = append(e.tab().Sequence, key)

		// Execution
		matches := e.Keymap.GetMatches(e.tab().Mode, e.tab().Sequence)
		if len(matches) == 1 && len(matches[0].Sequence) == len(e.tab().Sequence) {
			err := e.RunCommand(matches[0].Command)
			if err != nil {
				return err
			}
		} else if len(matches) == 0 {
			// if nothing matches just clear the sequence,
			// TODO:  perhaps even show an error...
			e.tab().Sequence = []humankey.InternalKey{}
			return nil
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
