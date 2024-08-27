package buffer

import (
	"github.com/Tesohh/femto/femath"
)

type Cursorer interface {
	Up(times int)
	Down(times int)
	Left(times int)
	Right(times int)
	ForceRight(times int)

	GoTo(pos femath.Vec2)
	Pos() femath.Vec2
}

type CRUDer interface {
	Insert(pos femath.Vec2, r rune)
	Replace(pos femath.Vec2, r rune) rune
	Get(pos femath.Vec2) rune
	Line() []rune
	Delete(pos femath.Vec2) rune
}

type Reader interface {
	Read() ([][]rune, error)
	Read1D() ([]rune, error)
}

type Writer interface {
	Write([][]rune) error
}

type Buffer interface {
	Cursorer
	CRUDer
	Reader
	Writer

	// io.Reader
	// io.Writer
	// io.Closer
}
