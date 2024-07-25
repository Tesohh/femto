package buffer

import "github.com/Tesohh/femto/femath"

type SliceLine []rune

// a Buffer implementation that simply uses go's slices.
type SliceBuffer struct {
	content []SliceLine
	Pos     femath.Vec2
}
