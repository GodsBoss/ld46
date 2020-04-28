package v2d

import (
	"math"
	"strconv"
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

func (v Vector) String() string {
	return "(" + strconv.FormatFloat(v.x, 'G', -1, 64) + ", " + strconv.FormatFloat(v.y, 'G', -1, 64) + ")"
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

func Abs(v Vector) Vector {
	if v.x < 0 {
		v.x = -v.x
	}
	if v.y < 0 {
		v.y = -v.y
	}
	return v
}

func Length(v Vector) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
