package structures

type Color struct {
	R, G, B, A float32
}

func NewRgbU8(r, g, b float32) Color {
	return Color{r, g, b, 1}
}

func NewRgbaU8(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}
