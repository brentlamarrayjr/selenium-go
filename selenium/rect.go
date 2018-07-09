package selenium

type Rect struct {
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

func (rect *Rect) Overlaps(rectangle *Rect) bool {
	return (rect.X < rectangle.X+rectangle.Width && rect.X+rect.Width > rectangle.X && rect.Y < rectangle.Y+rectangle.Height && rect.Y+rect.Height > rectangle.Y)
}
