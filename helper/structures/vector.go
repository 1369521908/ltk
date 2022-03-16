package structures

// Vector2 二维向量
type Vector2 struct {
	X float32
	Y float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func Vector3Zero() Vector3 {
	return Vector3{0, 0, 0}
}
