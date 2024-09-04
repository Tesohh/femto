package femath

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestRangeContainsVec2(t *testing.T) {
	assert.Equal(t, Range2{
		Start: Vec2{X: 5, Y: 0},
		End:   Vec2{X: 5, Y: 0},
	}.ContainsVec2(Vec2{X: 7, Y: 0}), false)

	assert.Equal(t, Range2{
		Start: Vec2{X: 0, Y: 0},
		End:   Vec2{X: 0, Y: 0},
	}.ContainsVec2(Vec2{X: 0, Y: 0}), true)

	assert.Equal(t, Range2{
		Start: Vec2{X: 0, Y: 0},
		End:   Vec2{X: 0, Y: 0},
	}.ContainsVec2(Vec2{X: 1, Y: 0}), false)

	assert.Equal(t, Range2{
		Start: Vec2{X: 0, Y: 0},
		End:   Vec2{X: 0, Y: 4},
	}.ContainsVec2(Vec2{X: 4, Y: 2}), true)

	assert.Equal(t, Range2{
		Start: Vec2{X: 2, Y: 0},
		End:   Vec2{X: 0, Y: 4},
	}.ContainsVec2(Vec2{X: 0, Y: 2}), true)

	assert.Equal(t, Range2{
		Start: Vec2{X: 2, Y: 0},
		End:   Vec2{X: 0, Y: 4},
	}.ContainsVec2(Vec2{X: 1, Y: 0}), false)

	assert.Equal(t, Range2{
		Start: Vec2{X: 2, Y: 0},
		End:   Vec2{X: 3, Y: 4},
	}.ContainsVec2(Vec2{X: 5, Y: 4}), false)
}
