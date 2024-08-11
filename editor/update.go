package editor

import (
	"strings"
	"unicode"

	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

func (e *Editor) Update() error {
	defer GracefulPanic(e)

	event := e.Screen.PollEvent()

	for _, p := range e.Plugins {
		switch p.(type) {
		case *DumbPlugin:
			break
		default:
			ev := p.Update(e, event)
			if ev != nil {
				event = ev
			}
		}
	}

	switch event := event.(type) {
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEsc && e.Win().Mode == "normal" {
			// clear sequence with esc in normal mode
			e.Win().Sequence = []humankey.InternalKey{}
			return nil
		}

		key := humankey.InternalKey{
			Key:     event.Key(),
			Rune:    event.Rune(),
			ModMask: event.Modifiers(),
		}

		if unicode.IsUpper(key.Rune) || strings.ContainsRune(humankey.UppercaseSpecialCharset, key.Rune) {
			// band aid solution for windows adding shift+ to uppercase chars
			key.ModMask &= ^tcell.ModShift
		}

		e.Win().Sequence = append(e.Win().Sequence, key)

		// Execution
		mode := e.Win().Mode
		if mode == "" { // failsafe
			mode = "normal"
		}

		winkeymap, err := e.Win().Keymap.ToInternal()
		if err != nil {
			return err
		}

		keymap := e.Keymap.MergeInternal(winkeymap)
		matches := keymap.GetMatches(mode, e.Win().Sequence)
		if len(matches) == 1 && len(matches[0].Sequence) == len(e.Win().Sequence) {
			err := e.RunCommand(matches[0].Command)
			e.Win().Sequence = []humankey.InternalKey{}
			if err != nil {
				return err
			}
		} else if len(matches) == 0 {
			// if nothing matches just clear the sequence,
			// TODO:  perhaps even show an error...

			pp := humankey.PrettyPrintSequence(e.Win().Sequence)
			e.Win().Sequence = []humankey.InternalKey{}
			return ErrNoKeyAssociated.Context(pp)
		}
	case *tcell.EventError:
		return event
	case *tcell.EventResize:
		w, h := event.Size()
		e.Screen.SetSize(w, h)
		e.Screen.Sync()
	case *EventCaught:
		return nil
	}

	return nil
}
