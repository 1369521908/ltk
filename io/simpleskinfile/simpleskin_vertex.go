package simpleskinfile

import (
	"bytes"
	"ltk/helper/structures"
)

// Vector2 二维向量
type Vector2 struct {
	X float32
	Y float32
}

// Vector3 三维顶点
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// SimpleSkinVertex skn
type SimpleSkinVertex struct {
	Position    Vector3          // 顶点坐标
	BoneIndices []byte           // 骨骼索引
	Weights     []float32        // 骨骼权重
	Normal      Vector3          // 法线
	UV          Vector2          // UV
	Color       structures.Color // 顶点颜色
}

func NewSimpleSkinVertex(br *bytes.Reader) *SimpleSkinVertex {
	v := &SimpleSkinVertex{}
	v.BoneIndices = make([]byte, 4)
	v.Weights = make([]float32, 4)
	return v
}
