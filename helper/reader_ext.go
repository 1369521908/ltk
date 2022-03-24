package helper

import (
	"bytes"
	"ltk/helper/structures"
)

func ReadColor(br *bytes.Reader, format structures.ColorFormat) structures.Color {

	color := structures.Color{}

	switch format {
	case structures.RgbU8:
		r, _ := br.ReadByte()
		g, _ := br.ReadByte()
		b, _ := br.ReadByte()
		R := float32(r / 255)
		G := float32(g / 255)
		B := float32(b / 255)
		return structures.NewRgbU8(R, G, B)
	case structures.RgbF32:
	case structures.RgbaU8:
		r, _ := br.ReadByte()
		g, _ := br.ReadByte()
		b, _ := br.ReadByte()
		a, _ := br.ReadByte()
		R := float32(r / 255)
		G := float32(g / 255)
		B := float32(b / 255)
		A := float32(a / 255)
		return structures.NewRgbaU8(R, G, B, A)
	case structures.RgbaF32:
	case structures.BgrU8:
	case structures.BgrF32:
	case structures.BgraU8:
	case structures.BgraF32:
	default:
	}
	return color
}
