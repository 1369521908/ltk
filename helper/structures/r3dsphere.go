package structures

import (
	"bytes"
	"unsafe"
)

type R3DSphere struct {
	Position Vector3
	Radius   float32
}

func NewR3DSphereByReader(br *bytes.Reader) *R3DSphere {
	r3DSphere := &R3DSphere{}

	position_ := make([]byte, 12)
	if _, err := br.Read(position_); err != nil {
		return nil
	}

	r3DSphere.Position.X = *(*float32)(unsafe.Pointer(&position_[0]))
	r3DSphere.Position.Y = *(*float32)(unsafe.Pointer(&position_[4]))
	r3DSphere.Position.Z = *(*float32)(unsafe.Pointer(&position_[8]))

	radius_ := make([]byte, 4)
	if _, err := br.Read(radius_); err != nil {
		return nil
	}
	r3DSphere.Radius = *(*float32)(unsafe.Pointer(&radius_[0]))

	return r3DSphere
}
