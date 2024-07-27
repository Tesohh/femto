package buffer

import (
	"testing"

	"github.com/Tesohh/femto/femath"
	"gotest.tools/v3/assert"
)

func TestMovement(t *testing.T) {
	buf := SliceBuffer{
		content: [][]rune{
			[]rune("func main() {"),
			[]rune("    fmt.Println(\"hey\")"),
			[]rune("    free()"),
			[]rune("}"),
		},
	}

	t.Run("vertical clamping", func(t *testing.T) {
		buf.GoTo(femath.Vec2{X: 21, Y: 1})
		buf.Down(1)

		assert.Equal(t, buf.Pos, femath.Vec2{X: 9, Y: 2})

		buf.Down(3)
		assert.Equal(t, buf.Pos, femath.Vec2{X: 0, Y: 3})
	})

	t.Run("horizontal clamping", func(t *testing.T) {
		buf.GoTo(femath.Vec2{X: 0, Y: 0})
		buf.Right(100)

		assert.Equal(t, buf.Pos, femath.Vec2{X: 12, Y: 0})
	})
}
