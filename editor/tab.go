package editor

import (
	"github.com/Tesohh/femto/buffer"
	"github.com/Tesohh/femto/humankey"
)

const (
	ModeNormal Mode = "normal"
	ModeInsert Mode = "insert"
)

type Mode string

type Tab struct {
	Buffer   buffer.Buffer
	FilePath string // set to "@@@scratchpad@@@" to make a scratchpad buffer 🤯
	Mode     string
	Sequence []humankey.InternalKey
}
