package skeletonfile

import "bufio"

type Skeleton struct {
	FORMAT_TOKEN int // 0x22FD4FC3; // FNV hash of the format token string
	Name         string
	IsLegacy     bool
	Joints       []SkeletonJoint
	Influences   []int16
	AssetName    string
}

func Read(br bufio.Reader, leaveOpen bool) {

}
