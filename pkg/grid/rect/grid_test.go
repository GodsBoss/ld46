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

func TestOffsets(t *testing.T) {
	testCases := map[string]struct {
		startField     rect.Field
		offset         rect.FieldOffset
		expectedColumn int
		expectedRow    int
	}{
		"left": {
			startField:     rect.CreateField(3, 1),
			offset:         rect.FieldOffsetLeft(),
			expectedColumn: 2,
			expectedRow:    1,
		},
		"right": {
			startField:     rect.CreateField(8, -2),
			offset:         rect.FieldOffsetRight(),
			expectedColumn: 9,
			expectedRow:    -2,
		},
		"up": {
			startField:     rect.CreateField(-5, 3),
			offset:         rect.FieldOffsetUp(),
			expectedColumn: -5,
			expectedRow:    2,
		},
		"down": {
			startField:     rect.CreateField(4, 3),
			offset:         rect.FieldOffsetDown(),
			expectedColumn: 4,
			expectedRow:    4,
		},
		"-2,+5": {
			startField:     rect.CreateField(3, -2),
			offset:         rect.FieldOffsetFromField(rect.CreateField(-2, 5)),
			expectedColumn: 1,
			expectedRow:    3,
		},
	}

	for name, testCase := range testCases {
		t.Run(
			name,
			func(t *testing.T) {
				resultField := testCase.offset.Apply(testCase.startField)
				if resultField.Column() != testCase.expectedColumn {
					t.Errorf("expected column %d, but got %d", testCase.expectedColumn, resultField.Column())
				}
				if resultField.Row() != testCase.expectedRow {
					t.Errorf("expected row %d, but got %d", testCase.expectedRow, resultField.Row())
				}
			},
		)
	}
}

func TestCoordinatesFromField(t *testing.T) {
	testCases := map[string]struct {
		fieldSize      v2d.Vector
		offset         v2d.Vector
		field          rect.Field
		expectedCoords v2d.Vector
	}{
		"zero_field_size_zero_offset": {
			field:          rect.CreateField(7, 18),
			expectedCoords: v2d.FromXY(0, 0),
		},
		"zero_field_size_with_offset": {
			offset:         v2d.FromXY(7.5, -2.125),
			field:          rect.CreateField(-100, 50),
			expectedCoords: v2d.FromXY(7.5, -2.125),
		},
		"with_field_size_zero_offset": {
			fieldSize:      v2d.FromXY(2.5, 7.5),
			field:          rect.CreateField(-4, 3),
			expectedCoords: v2d.FromXY(-10, 22.5),
		},
		"with_field_size_with_offset": {
			fieldSize:      v2d.FromXY(1.25, 8.5),
			offset:         v2d.FromXY(8.5, -20.125),
			field:          rect.CreateField(-5, 3),
			expectedCoords: v2d.FromXY(2.25, 5.375),
		},
	}

	for name, testCase := range testCases {
		t.Run(
			name,
			func(t *testing.T) {
				g := &rect.Grid{}
				g.Set(
					rect.FieldSize(testCase.fieldSize),
					rect.Offset(testCase.offset),
				)
				v := g.CoordinatesFromField(testCase.field)
				if v2d.Length(v2d.Sum(testCase.expectedCoords, v2d.Scale(v, -1))) > float64compareThreshold {
					t.Errorf("expected coordinates to be %s, but got %s", testCase.expectedCoords, v)
				}
			},
		)
	}
}

const float64compareThreshold = 0.000005
