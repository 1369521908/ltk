package skeletonfile

type SkeletonJoint struct {
	IsLegacy bool
	Flags    uint16
	ID       int16
	ParentID int16
	Radius   float32
	Name     string
}
