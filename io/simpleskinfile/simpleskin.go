package simpleskinfile

import (
	"bytes"
	"encoding/binary"
	"ltk/helper/structures"
	"ltk/logger"
	"sort"
)

type SimpleSkin struct {
	Submeshes []*SimpleSkinSubMesh
}

type SimpleSkinVertexType uint32

const (
	Basic SimpleSkinVertexType = iota
	Color
)

func NewSimpleSkin(data []byte, _leaveOpen bool) *SimpleSkin {
	skn := &SimpleSkin{}

	br := bytes.NewReader(data)

	magic_ := make([]byte, 4)
	if _, err := br.Read(magic_); err != nil {
		return nil
	}

	magic := binary.LittleEndian.Uint32(magic_)

	if magic != 0x00112233 {
		logger.Error("Invalid magic number in SimpleSkin file")
		return nil
	}

	major_ := make([]byte, 2)
	if _, err := br.Read(major_); err != nil {
		return nil
	}
	major := binary.LittleEndian.Uint16(major_)

	minor_ := make([]byte, 2)
	if _, err := br.Read(minor_); err != nil {
		return nil
	}
	minor := binary.LittleEndian.Uint16(minor_)

	if major != 0 && major != 2 && major != 4 && minor != 1 {
		logger.Error("Unsupported SimpleSkin version")
		return nil
	}

	indexCount := uint32(0)
	vertexCount := uint32(0)

	vertexType := Basic

	if major == 0 {
		indexCount_ := make([]byte, 4)
		if _, err := br.Read(indexCount_); err != nil {
			return nil
		}

		vertexCount_ := make([]byte, 4)
		if _, err := br.Read(vertexCount_); err != nil {
			return nil
		}
	} else {
		submeshCount_ := make([]byte, 4)
		if _, err := br.Read(submeshCount_); err != nil {
			return nil
		}

		submeshCount := binary.LittleEndian.Uint32(submeshCount_)
		for i := uint32(0); i < submeshCount; i++ {
			skn.Submeshes = append(skn.Submeshes, NewSimpleSkinSubMesh(br))
		}
		if major == 4 {
			flags := make([]byte, 4)
			if _, err := br.Read(flags); err != nil {
				return nil
			}
		}

		indexCount_ := make([]byte, 4)
		if _, err := br.Read(indexCount_); err != nil {
			return nil
		}
		indexCount = binary.LittleEndian.Uint32(indexCount_)

		vertexCount_ := make([]byte, 4)
		if _, err := br.Read(vertexCount_); err != nil {
			return nil
		}
		vertexCount = binary.LittleEndian.Uint32(vertexCount_)

		vertexSize := uint32(52)
		if major == 4 {
			vertexSize_ := make([]byte, 4)
			if _, err := br.Read(vertexSize_); err != nil {
				return nil
			}
			vertexSize = binary.LittleEndian.Uint32(vertexSize_)
		}

		vertexType = Basic
		var boundingBox *structures.R3DBox
		var boundingSphere *structures.R3DSphere
		if major == 4 {
			vertexType_ := make([]byte, 4)
			if _, err := br.Read(vertexType_); err != nil {
				return nil
			}
			vertexType = SimpleSkinVertexType(binary.LittleEndian.Uint32(vertexType_))
			boundingBox = structures.NewR3DBoxWithReader(br)
			boundingSphere = structures.NewR3DSphereByReader(br)
		} else {
			boundingBox = new(structures.R3DBox)
			boundingSphere = new(structures.R3DSphere)
		}

		if vertexSize == 0 {

		}

		if boundingBox == nil {

		}

		if boundingSphere == nil {

		}

	}

	indices := make([]uint16, 0)

	vertices := make([]*SimpleSkinVertex, 0)

	for i := uint32(0); i < indexCount; i++ {
		index_ := make([]byte, 2)
		if _, err := br.Read(index_); err != nil {
			return nil
		}
		index := binary.LittleEndian.Uint16(index_)
		indices = append(indices, index)
	}

	for i := uint32(0); i < vertexCount; i++ {
		vertices = append(vertices, NewSimpleSkinVertex(br, vertexType))
	}

	if major == 0 {
		skn.Submeshes = append(skn.Submeshes, NewSimpleSkinSubMeshByName("Base", indices, vertices))
	} else {
		for _, submesh := range skn.Submeshes {
			submeshIndices := indices[submesh._startIndex : submesh._startIndex+submesh._indexCount]
			indices := make([]uint16, len(submeshIndices))
			copy(indices, submeshIndices)
			sort.Slice(indices, func(i, j int) bool {
				return indices[i] < indices[j]
			})
			minIndex := indices[0]

			indices = make([]uint16, 0)
			for _, index := range submeshIndices {
				submesh.Indices = append(submesh.Indices, index-minIndex)
			}
			submesh.Vertices = vertices[submesh._startVertex : submesh._startVertex+submesh._vertexCount]
		}
	}

	return skn
}
