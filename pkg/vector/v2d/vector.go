package v2d

import (
	"math"
)

type Vector struct {
	x float64
	y float64
}

func FromXY(x, y float64) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

func Zero() Vector {
	return Vector{}
}

func UnitX() Vector {
	return Vector{
		x: 1,
	}
}

func UnitY() Vector {
	return Vector{
		y: 1,
	}
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}

func (v Vector) XY() (float64, float64) {
	return v.x, v.y
}

func Sum(vs ...Vector) Vector {
	result := Vector{}
	for i := range vs {
		result.x, result.y = result.x+vs[i].x, result.y+vs[i].y
	}
	return result
}

func Scale(v Vector, factor float64) Vector {
	return Vector{
		x: v.x * factor,
		y: v.y * factor,
	}
}

func Length(v Vector) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
