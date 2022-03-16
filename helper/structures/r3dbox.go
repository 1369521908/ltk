package structures

import (
	"bytes"
	"unsafe"
)

// TODO tbd
type R3DBox struct {
	Min Vector3
	Max Vector3
}

func NewR3DBoxWithReader(br *bytes.Reader) *R3DBox {
	r3DBox := &R3DBox{}

	// float32 => 4*byte and count is 3
	min_ := make([]byte, 12)
	if _, err := br.Read(min_); err != nil {
		return nil
	}
	r3DBox.Min.X = *(*float32)(unsafe.Pointer(&min_[0]))
	r3DBox.Min.Y = *(*float32)(unsafe.Pointer(&min_[4]))
	r3DBox.Min.Z = *(*float32)(unsafe.Pointer(&min_[8]))

	max_ := make([]byte, 12)
	if _, err := br.Read(max_); err != nil {
		return nil
	}
	r3DBox.Max.X = *(*float32)(unsafe.Pointer(&max_[0]))
	r3DBox.Max.Y = *(*float32)(unsafe.Pointer(&max_[4]))
	r3DBox.Max.Z = *(*float32)(unsafe.Pointer(&max_[8]))

	return r3DBox
}
