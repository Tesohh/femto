package buffer

import "github.com/Tesohh/femto/femath"

func (s *SliceBuffer) Down(times int) {
	s.pos.Y += times
	s.pos.Y = femath.Clamp(s.pos.Y, 0, len(s.content)-1)
	s.pos.X = femath.Clamp(s.pos.X, 0, len(s.content[s.pos.Y]))
}

func (s *SliceBuffer) Up(times int) {
	s.Down(times * -1)
}

func (s *SliceBuffer) Right(times int) {
	s.pos.X += times
	s.pos.X = femath.Clamp(s.pos.X, 0, len(s.content[s.pos.Y])-1)
}

// Just like Right but lets you go to the end line
func (s *SliceBuffer) ForceRight(times int) {
	s.pos.X += times
	s.pos.X = femath.Clamp(s.pos.X, 0, len(s.content[s.pos.Y]))
}

func (s *SliceBuffer) Left(times int) {
	s.Right(times * -1)
}

func (s *SliceBuffer) GoTo(pos femath.Vec2) {
	s.pos.Y = femath.Clamp(pos.Y, 0, len(s.content)-1)
	s.pos.X = femath.Clamp(pos.X, 0, len(s.content[s.pos.Y])-1)
}

func (s *SliceBuffer) Pos() femath.Vec2 {
	return s.pos
}
