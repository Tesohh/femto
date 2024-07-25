package buffer

import (
	"github.com/Tesohh/femto/femath"
)

type Cursorer interface {
	Up(times int)
	Down(times int)
	Left(times int)
	Right(times int)

	GoTo(pos femath.Vec2)
}

type CRUDer interface {
	Insert(pos femath.Vec2, r rune)
	Replace(pos femath.Vec2, r rune) rune
	Get(pos femath.Vec2) rune
	Delete(pos femath.Vec2) rune
}

type Buffer interface {
	Cursorer
	CRUDer

	// io.Reader
	// io.Writer
	// io.Closer
}
