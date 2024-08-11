package editor

type Tab struct {
	Windows            []Window
	FocusedWindowIndex int
}

func (t *Tab) RegisterWindow(w Window) *Window {
	setupWindow(&w)
	t.Windows = append(t.Windows, w)
	return &t.Windows[len(t.Windows)-1]
}

// Side effect: this also sets editor's FocusedWindowIndex to -1 if a window is found
func (t *Tab) FocusWindow(e *Editor, id string) error {
	for i, w := range t.Windows {
		if w.Id == id {
			e.FocusedWindowIndex = -1
			t.FocusedWindowIndex = i
			break
		}
	}

	// if we get to this point it means we found nothing
	return ErrNoWindowFoundForId.Context(id)
}

func (t *Tab) GetWindow(id string) (*Window, error) {
	for _, w := range t.Windows {
		if w.Id == id {
			return &w, nil
		}
	}
	return nil, ErrNoWindowFoundForId.Context(id)
}
