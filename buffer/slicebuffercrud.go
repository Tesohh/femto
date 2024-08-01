package buffer

import (
	"slices"

	"github.com/Tesohh/femto/femath"
)

func (s *SliceBuffer) Insert(pos femath.Vec2, r rune) {
	s.content[pos.Y] = slices.Insert(s.content[pos.Y], pos.X, r)
}

func (s *SliceBuffer) Replace(pos femath.Vec2, r rune) rune {
	old := s.content[pos.Y][pos.X]
	s.content[pos.Y][pos.X] = r
	return old
}

func (s *SliceBuffer) Get(pos femath.Vec2) rune {
	return s.content[pos.Y][pos.X]
}

func (s *SliceBuffer) Line() []rune {
	return s.content[s.pos.Y]
}

func (s *SliceBuffer) Delete(pos femath.Vec2) rune {
	old := s.content[pos.Y][pos.X]
	s.content[pos.Y] = slices.Delete(s.content[pos.Y], pos.X, pos.X+1)
	return old
}
