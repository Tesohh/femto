package editor

import "github.com/Tesohh/femto/buffer"

const (
	ModeNormal = "normal"
	ModeInsert = "insert"
)

type Tab struct {
	Buffer   buffer.Buffer
	FilePath string // set to "@@@scratchpad@@@" to make a scratchpad buffer 🤯
	Mode     string
}
