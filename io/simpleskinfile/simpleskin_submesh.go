package simpleskinfile

import (
	"bytes"
	"encoding/binary"
	"strings"
)

// SimpleSkinSubMesh 子网格
type SimpleSkinSubMesh struct {
	Name         string
	Vertices     []*SimpleSkinVertex
	Indices      []uint16
	_startVertex uint32
	_vertexCount uint32
	_startIndex  uint32
	_indexCount  uint32
}

func NewSimpleSkinSubMeshByName(name string, indices []uint16, vertices []*SimpleSkinVertex) *SimpleSkinSubMesh {
	return &SimpleSkinSubMesh{
		Name:         name,
		Vertices:     vertices,
		Indices:      indices,
		_startVertex: 0,
		_vertexCount: uint32(len(vertices)),
		_startIndex:  0,
		_indexCount:  uint32(len(indices)),
	}
}

func NewSimpleSkinSubMesh(br *bytes.Reader) *SimpleSkinSubMesh {

	sub := &SimpleSkinSubMesh{}

	name_ := make([]byte, 64)
	if _, err := br.Read(name_); err != nil {
		return nil
	}
	subName := string(name_)
	sub.Name = strings.ReplaceAll(subName, `\0`, ``)

	_startVertex_ := make([]byte, 4)
	if _, err := br.Read(_startVertex_); err != nil {
		return nil
	}
	sub._startVertex = binary.LittleEndian.Uint32(_startVertex_)

	_vertexCount_ := make([]byte, 4)
	if _, err := br.Read(_vertexCount_); err != nil {
		return nil
	}
	sub._vertexCount = binary.LittleEndian.Uint32(_vertexCount_)

	_startIndex_ := make([]byte, 4)
	if _, err := br.Read(_startIndex_); err != nil {
		return nil
	}
	sub._startIndex = binary.LittleEndian.Uint32(_startIndex_)

	_indexCount_ := make([]byte, 4)
	if _, err := br.Read(_indexCount_); err != nil {
		return nil
	}
	sub._indexCount = binary.LittleEndian.Uint32(_indexCount_)

	return sub
}
