package editor

import "github.com/Tesohh/femto/buffer"

const (
	ModeNormal Mode = "normal"
	ModeInsert Mode = "insert"
)

type Mode string

type Tab struct {
	Buffer   buffer.Buffer
	FilePath string // set to "@@@scratchpad@@@" to make a scratchpad buffer 🤯
	Mode     Mode
}
