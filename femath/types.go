package femath

type Vec2 struct {
	X int
	Y int
}

type Range2 struct {
	Start Vec2
	End   Vec2
}

func (r Range2) ContainsVec2(v Vec2) bool {
	start, end := r.Start, r.End
	if (start.Y > end.Y) || (start.Y == end.Y && start.X > end.X) {
		start, end = end, start
	}

	if v.Y > start.Y && v.Y < end.Y {
		return true
	}

	if v.Y == start.Y {
		return v.X >= start.X
	} else if v.Y == end.Y {
		return v.X <= end.X
	}

	return false
}
