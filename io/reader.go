package io

import (
	"bytes"
	"encoding/binary"
)

type CsReader bytes.Reader

func (r *CsReader) Read(p []byte) (n int, err error) {
	return (*bytes.Reader)(r).Read(p)
}

func (r *CsReader) ReadUint16() (uint uint16, err error) {
	b := make([]byte, 2)
	if _, err := r.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b), nil
}

func (r *CsReader) ReadUint32() (uint uint32, err error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func (r *CsReader) ReadUint64() (uint uint64, err error) {
	b := make([]byte, 8)
	if _, err := r.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b), nil
}