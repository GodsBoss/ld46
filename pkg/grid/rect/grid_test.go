package rect_test

import (
	"fmt"

	"github.com/GodsBoss/ld46/pkg/grid/rect"
	"github.com/GodsBoss/ld46/pkg/vector/v2d"

	"testing"
)

func TestGridSetDoesNotApplyOptionsBeforeAnErrorneousOption(t *testing.T) {
	g := &rect.Grid{}
	g.Set(
		rect.FieldSize(v2d.FromXY(100.0, 100.0)),
		rect.Size(-1, 3),
	)
	f, err := g.FieldFromCoordinates(v2d.FromXY(0.0, 0.0))
	if f != nil {
		t.Errorf("expected no field, but got %+v", f)
	}
	if err == nil {
		t.Errorf("expected error, got none")
	}
}

func TestOptionErrors(t *testing.T) {
	testCases := map[string]struct {
		option rect.GridOption
	}{
		"columns_lesser_than_zero": {
			option: rect.Size(-2, 3),
		},
		"rows_lesser_than_zero": {
			option: rect.Size(4, -1),
		},
		"field_width_lesser_than_zero": {
			option: rect.FieldSize(v2d.FromXY(-5.0, 8.0)),
		},
		"field_height_lesser_than_zero": {
			option: rect.FieldSize(v2d.FromXY(2.5, -7.5)),
		},
	}
	for name := range testCases {
		option := testCases[name].option
		t.Run(
			name,
			func(t *testing.T) {
				g := &rect.Grid{}
				err := g.Set(option)
				if err == nil {
					t.Errorf("missing error for invalid option")
				}
			},
		)
	}
}

func TestGridContains(t *testing.T) {
	g := &rect.Grid{}
	g.Set(rect.Size(8, 6))
	testCases := []struct {
		column    int
		row       int
		contained bool
	}{
		{0, 0, true},
		{-1, 2, false},
		{2, -1, false},
		{7, 5, true},
		{8, 3, false},
		{7, 6, false},
		{3, 3, true},
	}
	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("(%d, %d) contained? %t", testCase.column, testCase.row, testCase.contained),
			func(t *testing.T) {
				field := rect.CreateField(testCase.column, testCase.row)
				if testCase.contained && !g.Contains(field) {
					t.Errorf("grid should contain field, but does not")
				}
				if !testCase.contained && g.Contains(field) {
					t.Errorf("grid shouldn't contain field, but does")
				}
			},
		)
	}
}

func TestGridFieldFromCoordinates(t *testing.T) {
	grid := &rect.Grid{}
	grid.Set(
		rect.Offset(v2d.FromXY(-150.0, 75)),
		rect.FieldSize(v2d.FromXY(40, 20)),
	)
	testCases := []struct {
		x   float64
		y   float64
		col int
		row int
	}{
		{
			x:   -151.0,
			y:   74.0,
			col: -1,
			row: -1,
		},
		{
			x:   10.0,
			y:   115,
			col: 4,
			row: 2,
		},
	}
	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("%f,%f", testCase.x, testCase.y),
			func(t *testing.T) {
				field, err := grid.FieldFromCoordinates(v2d.FromXY(testCase.x, testCase.y))
				if err != nil {
					t.Fatalf("expected no error, but got %+v", err)
				}
				if field.Column() != testCase.col {
					t.Errorf("expected field column to be %d, but got %d", testCase.col, field.Column())
				}
				if field.Row() != testCase.row {
					t.Errorf("expected field row to be %d, but got %d", testCase.row, field.Row())
				}
			},
		)
	}
}
