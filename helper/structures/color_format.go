package structures

type ColorFormat int

const (
	RgbU8 ColorFormat = iota
	RgbF32
	RgbaU8
	RgbaF32
	BgrU8
	BgrF32
	BgraU8
	BgraF32
)
