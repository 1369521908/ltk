package simpleskinfile

import (
	"bytes"
	"ltk/helper/structures"
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

func NewSimpleSkinVertex(br *bytes.Reader) *SimpleSkinVertex {
	v := &SimpleSkinVertex{}
	v.BoneIndices = make([]byte, 4)
	v.Weights = make([]float32, 4)
	return v
}
