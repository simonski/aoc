package utils

type Point2D struct {
	X int
	Y int
}

type Point2DF struct {
	X float64
	Y float64
}

func (point *Point2DF) Translate(x float64, y float64) *Point2DF {
	copy := Point2DF{X: point.X + x, Y: point.Y + y}
	return &copy
}

type Point3D struct {
	X int
	Y int
	Z int
}

type Point4D struct {
	X int
	Y int
	Z int
	W int
}

// Rotates this point around origin 0, 0
func (p *Point2D) Rotate(degrees int) {
	origin := &Point2D{0, 0}
	p.RotateAroundOrigin(degrees, origin)
}

// RotatesAroundOrigin rotates this point around the specified origin
func (p *Point2D) RotateAroundOrigin(degrees int, origin *Point2D) {

	if degrees < 0 {
		degrees = 360 + degrees
	}

	x_original := p.X - origin.X
	y_original := p.Y - origin.Y

	x := 0
	y := 0

	if degrees == 90 {
		// 90 cw (y, -x)
		x = y_original
		y = -x_original
	} else if degrees == 180 {
		// 180 cw (x,y) -> (-x, -y)
		x = -x_original
		y = -y_original
	} else if degrees == 270 {
		// 180 cw (x,y) -> (-y, x)
		x = -y_original
		y = x_original
	}

	x += origin.X
	y += origin.Y

	p.X = x
	p.Y = y

}
