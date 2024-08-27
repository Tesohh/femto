package buffer

import (
	"fmt"
)

func (s *SliceBuffer) Read() ([][]rune, error) {
	return [][]rune(s.content), nil
}

func (s *SliceBuffer) Read1D() ([]rune, error) {
	r := []rune{}
	for _, line := range s.content {
		for _, char := range line {
			r = append(r, char)
		}
		r = append(r, '\n')
	}

	return r, nil
}

func (s *SliceBuffer) Write(r [][]rune) error {
	if r == nil {
		return fmt.Errorf("content is nil")
	}
	s.content = r
	return nil
}
