package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (e *Editor) Draw(screen *ebiten.Image) {
	x := 0
	y := 32

	e.TextRenderer.SetTarget(screen)

	buf, err := e.tab().Buffer.Read()
	if err != nil {
		panic(err)
	}

	for _, line := range buf {
		for _, char := range line {
			s := string(char)
			e.TextRenderer.Draw(s, x, y)
			rect := e.TextRenderer.SelectionRect(s)
			x += rect.Width.Ceil()
		}
		rect := e.TextRenderer.SelectionRect("a")
		y += rect.Height.Ceil()
	}
}
