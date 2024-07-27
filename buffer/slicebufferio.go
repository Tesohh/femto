package buffer

import "fmt"

func (s *SliceBuffer) Read() ([][]rune, error) {
	return [][]rune(s.content), nil
}

func (s *SliceBuffer) Write(r [][]rune) error {
	if r == nil {
		return fmt.Errorf("content is nil")
	}
	s.content = r
	return nil
}
