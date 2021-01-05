package main

type Point2D struct {
	x int
	y int
}

type Point3D struct {
	x int
	y int
	z int
}

type Point4D struct {
	x int
	y int
	z int
	w int
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

	x_original := p.x - origin.x
	y_original := p.y - origin.y

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

	x += origin.x
	y += origin.y

	p.x = x
	p.y = y

}
