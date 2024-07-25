package buffer

import "github.com/Tesohh/femto/femath"

func (s *SliceBuffer) Down(times int) {
	s.Pos.Y += times
	s.Pos.Y = femath.Clamp(s.Pos.Y, 0, len(s.content)-1)
	s.Pos.X = femath.Clamp(s.Pos.X, 0, len(s.content[s.Pos.Y])-1)
}

func (s *SliceBuffer) Up(times int) {
	s.Down(times * -1)
}

func (s *SliceBuffer) Right(times int) {
	s.Pos.X += times
	s.Pos.X = femath.Clamp(s.Pos.X, 0, len(s.content[s.Pos.Y])-1)
}

func (s *SliceBuffer) Left(times int) {
	s.Right(times * -1)
}

func (s *SliceBuffer) GoTo(pos femath.Vec2) {
	s.Pos.Y = femath.Clamp(pos.Y, 0, len(s.content)-1)
	s.Pos.X = femath.Clamp(pos.X, 0, len(s.content[s.Pos.Y])-1)
}
