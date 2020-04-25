package v2d_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/GodsBoss/ld46/pkg/vector/v2d"
)

func TestVector(t *testing.T) {
	assertEqual(t, "zero", 0, 0, v2d.Zero())
	assertEqual(t, "x unit", 1, 0, v2d.UnitX())
	assertEqual(t, "y unit", 0, 1, v2d.UnitY())
	assertEqual(t, "v", -17.0, 5.0, v2d.FromXY(-17.0, 5.0))
	assertEqual(t, "scaled", 3.0, -6.0, v2d.Scale(v2d.FromXY(-1.0, 2.0), -3.0))
	assertEqual(t, "sum", -3.0, 4.0, v2d.Sum(v2d.UnitX(), v2d.UnitY(), v2d.FromXY(-4.0, 3.0), v2d.Zero()))
}

func TestLength(t *testing.T) {
	l := v2d.Length(v2d.FromXY(3.0, -4.0))
	diff := math.Abs(l - 5.0)
	if diff > 0.0000001 {
		t.Errorf("expected length of %s to be %f, but got %f", formatXY(3.0, -4.0), 5.0, l)
	}
}

func assertEqual(t *testing.T, name string, x, y float64, actual v2d.Vector) {
	if x != actual.X() || y != actual.Y() {
		t.Errorf("expected %s to be %s, but got %s.", name, formatXY(x, y), formatVector(actual))
	}
	actualX, actualY := actual.XY()
	if x != actualX || y != actualY {
		t.Errorf("expected %s to be %s, but got %s.", name, formatXY(x, y), formatVector(actual))
	}
}

func formatXY(x, y float64) string {
	return fmt.Sprintf("(%f, %f)", x, y)
}

func formatVector(v v2d.Vector) string {
	return formatXY(v.X(), v.Y())
}
