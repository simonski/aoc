package d18

type Point3D struct {
	x int
	y int
	z int
}

func NewPoint3D(x int, y int, z int) *Point3D {
	return &Point3D{x: x, y: y, z: z}
}
