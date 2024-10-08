package buffer

import (
	"testing"

	"github.com/Tesohh/femto/femath"
	"gotest.tools/v3/assert"
)

func TestCRUD(t *testing.T) {
	buf := SliceBuffer{
		content: [][]rune{
			[]rune("fun main() {"),
		},
	}

	t.Run("inserting", func(t *testing.T) {
		buf.Insert(femath.Vec2{X: 3, Y: 0}, 'c')
		assert.DeepEqual(t, buf.content[0], []rune("func main() {"))
	})

	t.Run("bulk inserting", func(t *testing.T) {
		s := "[]string args"
		buf.pos.X = 10
		for _, v := range s {
			buf.Insert(buf.pos, v)
			buf.pos.X += 1
		}

		assert.DeepEqual(t, buf.content[0], []rune("func main([]string args) {"))
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
