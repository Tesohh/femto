package editor

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type EventCaught struct {
	when time.Time
}

func (c *EventCaught) When() time.Time {
	return c.when
}

type CommandBarEvent struct {
	Msg   string
	Style tcell.Style
	Time  time.Time
}

func (c *CommandBarEvent) When() time.Time {
	return c.Time
}
