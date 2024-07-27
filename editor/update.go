package editor

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Editor) Update() error {
	keys := []rune{}
	keys = ebiten.IsKeyPressed()

	for _, v := range keys {
		fmt.Println(string(v), v)
	}

	return nil
}
