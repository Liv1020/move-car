package components

// Page Page
type Page struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// GetOffset GetOffset
func (t *Page) GetOffset() int {
	if t.Page == 0 {
		return 0
	}
	return (t.Page - 1) * t.Size
}

// GetLimit GetLimit
func (t *Page) GetLimit() int {
	if t.Size == 0 {
		return 15
	}
	return t.Size
}
