package buffer

import "github.com/Tesohh/femto/femath"

// a Buffer implementation that simply uses go's slices.
type SliceBuffer struct {
	content [][]rune
	pos     femath.Vec2
}
