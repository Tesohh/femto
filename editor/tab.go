package editor

type Tab struct {
	Windows            []Window
	FocusedWindowIndex int
}

func (t *Tab) RegisterWindow(w Window) {
	setupWindow(&w)
	t.Windows = append(t.Windows, w)
}
