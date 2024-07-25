package buffer

import (
	"testing"

	"github.com/Tesohh/femto/femath"
	"gotest.tools/v3/assert"
)

func TestMovement(t *testing.T) {
	buf := SliceBuffer{
		content: []SliceLine{
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

func TestCRUD(t *testing.T) {
	buf := SliceBuffer{
		content: []SliceLine{
			[]rune("fun main() {"),
		},
	}

	t.Run("inserting", func(t *testing.T) {
		buf.Insert(femath.Vec2{X: 3, Y: 0}, 'c')
		assert.DeepEqual(t, buf.content[0], SliceLine("func main() {"))
	})

	t.Run("bulk inserting", func(t *testing.T) {
		s := "[]string args"
		buf.Pos.X = 10
		for _, v := range s {
			buf.Insert(buf.Pos, v)
			buf.Pos.X += 1
		}

		assert.DeepEqual(t, buf.content[0], SliceLine("func main([]string args) {"))
	})

	t.Run("getting", func(t *testing.T) {
		r := buf.Get(femath.Vec2{})
		assert.Equal(t, r, 'f')
	})

	t.Run("replacing", func(t *testing.T) {
		r := buf.Replace(femath.Vec2{}, 'p')

		assert.Equal(t, r, 'f')
		assert.Equal(t, buf.Get(femath.Vec2{}), 'p')
	})

	n := len(buf.content[0])
	t.Run("deleting", func(t *testing.T) {
		r := buf.Delete(femath.Vec2{})

		assert.Equal(t, r, 'p')
		assert.Equal(t, len(buf.content[0]), n-1)
	})
}
