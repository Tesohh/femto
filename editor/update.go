package editor

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
	"unicode"

	"github.com/Tesohh/femto/humankey"
	"github.com/gdamore/tcell/v2"
)

type EventCaught struct {
	when time.Time
}

func (c *EventCaught) When() time.Time {
	return c.when
}

func (e *Editor) Update() error {
	event := e.Screen.PollEvent()
	slog.Info(fmt.Sprintf("got new event: %#v", event))

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
		if event.Key() == tcell.KeyEsc && e.Tab().Mode == "normal" {
			// clear sequence with esc in normal mode
			e.Tab().Sequence = []humankey.InternalKey{}
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

		e.Tab().Sequence = append(e.Tab().Sequence, key)

		// Execution
		matches := e.Keymap.GetMatches(e.Tab().Mode, e.Tab().Sequence)
		if len(matches) == 1 && len(matches[0].Sequence) == len(e.Tab().Sequence) {
			err := e.RunCommand(matches[0].Command)
			e.Tab().Sequence = []humankey.InternalKey{}
			if err != nil {
				return err
			}
		} else if len(matches) == 0 {
			// if nothing matches just clear the sequence,
			// TODO:  perhaps even show an error...

			pp := humankey.PrettyPrintSequence(e.Tab().Sequence)
			e.Tab().Sequence = []humankey.InternalKey{}
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
