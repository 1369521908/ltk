package helper

import (
	"bufio"
	"encoding/binary"
)

type CsReader struct {
	Br *bufio.Reader
}

func (r *CsReader) ReadUint16() (uint uint16, err error) {
	b := make([]byte, 2)
	if _, err := r.Br.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b), nil
}

func (r *CsReader) ReadUint32() (uint uint32, err error) {
	b := make([]byte, 4)
	if _, err := r.Br.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func (r *CsReader) ReadUint64() (uint uint64, err error) {
	b := make([]byte, 8)
	if _, err := r.Br.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b), nil
}
