package yic

import (
	"fmt"
)

type vector2D struct {
	X int
	Y int
}

func (v vector2D) scale(f int) vector2D {
	v.X *= f
	v.Y *= f
	return v
}

func (v vector2D) abs() vector2D {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	return v
}

func (v vector2D) sum() int {
	return v.X + v.Y
}

func (v vector2D) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func addVector2Ds(vs ...vector2D) vector2D {
	v := vector2D{}
	for i := range vs {
		v.X += vs[i].X
		v.Y += vs[i].Y
	}
	return v
}

var directionUp = vector2D{
	X: 0,
	Y: -1,
}

var directionDown = vector2D{
	X: 0,
	Y: 1,
}

var directionLeft = vector2D{
	X: -1,
	Y: 0,
}

var directionRight = vector2D{
	X: 1,
	Y: 0,
}
