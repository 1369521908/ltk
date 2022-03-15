package structures

type Color struct {
	R, G, B, A float32
}

func NewColor() *Color {
	return &Color{1, 1, 1, 1}
}
