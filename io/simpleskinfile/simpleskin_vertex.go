package simpleskinfile

import (
	"bytes"
	"ltk/helper/structures"
	"unsafe"
)

// SimpleSkinVertex skn
type SimpleSkinVertex struct {
	Position    structures.Vector3 // 顶点坐标
	BoneIndices []byte             // 骨骼索引
	Weights     []float32          // 骨骼权重
	Normal      structures.Vector3 // 法线
	UV          structures.Vector2 // UV
	Color       structures.Color   // 顶点颜色
}

func NewSimpleSkinVertex(br *bytes.Reader, vertexType SimpleSkinVertexType) *SimpleSkinVertex {
	v := &SimpleSkinVertex{}

	position := make([]byte, 4*3)
	if _, err := br.Read(position); err != nil {
		return nil
	}
	v.Position.X = *(*float32)(unsafe.Pointer(&position[0]))
	v.Position.Y = *(*float32)(unsafe.Pointer(&position[4]))
	v.Position.Z = *(*float32)(unsafe.Pointer(&position[8]))

	v.BoneIndices = make([]byte, 4)
	if _, err := br.Read(v.BoneIndices); err != nil {
		return nil
	}

	weights := make([]byte, 4*4)
	if _, err := br.Read(weights); err != nil {
		return nil
	}

	sizeof := unsafe.Sizeof(float32(0))
	if sizeof == 0 {
		return nil
	}
	v.Weights = make([]float32, 4)
	for i := 0; i < 4; i++ {
		index := i * int(sizeof)
		v.Weights[i] = *(*float32)(unsafe.Pointer(&weights[index]))
	}

	normal := make([]byte, 4*3)
	if _, err := br.Read(normal); err != nil {
		return nil
	}
	v.Normal.X = *(*float32)(unsafe.Pointer(&normal[0]))
	v.Normal.Y = *(*float32)(unsafe.Pointer(&normal[4]))
	v.Normal.Z = *(*float32)(unsafe.Pointer(&normal[8]))

	uv := make([]byte, 4*2)
	if _, err := br.Read(uv); err != nil {
		return nil
	}
	v.UV.X = *(*float32)(unsafe.Pointer(&uv[0]))
	v.UV.Y = *(*float32)(unsafe.Pointer(&uv[4]))

	if vertexType == Color {
		color := make([]byte, 4*4)
		if _, err := br.Read(color); err != nil {
			return nil
		}
		v.Color.R = *(*float32)(unsafe.Pointer(&color[0]))
		v.Color.G = *(*float32)(unsafe.Pointer(&color[4]))
		v.Color.B = *(*float32)(unsafe.Pointer(&color[8]))
		v.Color.A = *(*float32)(unsafe.Pointer(&color[12]))
	}

	return v
}
