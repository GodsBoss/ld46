package rect

import (
	"errors"
	"math"
	"strconv"

	"github.com/GodsBoss/ld46/pkg/vector/v2d"
)

// Grid represents a finite rectangular grid, providing methods for calculating
// field indexes and more. The zero value is usable (although probably not very useful).
type Grid struct {
	columns int
	rows    int

	offset    v2d.Vector
	fieldSize v2d.Vector
}

// Set applies options to the grid. If any option fails and causes an error,
// the grid remains unchanged and that error is returned. No more options are
// applied in this case.
// Unlike grid methods for queries and calcuations, Set is not safe for concurrency.
// The typical use case is: Set the options on the grid once, then use methods
// like FieldFromCoordinates from as many Go routines as you like.
func (grid *Grid) Set(options ...GridOption) error {
	tempGrid := *grid
	for i := range options {
		if err := options[i].applyTo(&tempGrid); err != nil {
			return err
		}
	}
	*grid = tempGrid
	return nil
}

// GridOption is used by Grid.Set().
type GridOption interface {
	applyTo(*Grid) error
}

type gridOptionFunc func(*Grid) error

func (f gridOptionFunc) applyTo(grid *Grid) error {
	return f(grid)
}

// Size is an option for setting the size of the grid, i.e. how many columns and rows the grid has.
func Size(columns int, rows int) GridOption {
	return gridOptionFunc(
		func(grid *Grid) error {
			if columns < 0 || rows < 0 {
				return errors.New("columns and height must both be >= 0, but got (" + strconv.Itoa(columns) + ", " + strconv.Itoa(rows) + ")")
			}
			grid.columns = columns
			grid.rows = rows
			return nil
		},
	)
}

// Offset is an option for setting the offset of the grid.
func Offset(offset v2d.Vector) GridOption {
	return gridOptionFunc(
		func(grid *Grid) error {
			grid.offset = offset
			return nil
		},
	)
}

// FieldSize is an option for setting the size of a single field on the grid.
func FieldSize(fieldSize v2d.Vector) GridOption {
	return gridOptionFunc(
		func(grid *Grid) error {
			if fieldSize.X() < 0 || fieldSize.Y() < 0 {
				return errors.New("field size must be positive, but got " + fieldSize.String())
			}
			grid.fieldSize = fieldSize
			return nil
		},
	)
}

// Contains checks wether a field is inside the grid.
func (grid *Grid) Contains(field Field) bool {
	return field.column >= 0 && field.row >= 0 && field.column < grid.columns && field.row < grid.rows
}

// FieldFromCoordinates takes a vector and returns its corresponding grid field. This field is not
// necessarely contained within the grid, this can be checked with Grid.Contains().
// Returns an error and no field if the field size in any dimension is zero.
func (grid *Grid) FieldFromCoordinates(pos v2d.Vector) (*Field, error) {
	if grid.fieldSize.X() <= 0 || grid.fieldSize.Y() <= 0 {
		return nil, errors.New("grid size must be > 0 both horizontally and vertically to calculate a field from coordinates, but is " + grid.fieldSize.String())
	}
	pos = v2d.Sum(pos, v2d.Scale(grid.offset, -1.0))
	return &Field{
		column: int(math.Floor(pos.X() / grid.fieldSize.X())),
		row:    int(math.Floor(pos.Y() / grid.fieldSize.Y())),
	}, nil
}

// Field is a single field of a grid, identified by column and row. Fields are value objects, i.e. they never change.
// Fields can safely be used as map keys.
type Field struct {
	column int
	row    int
}

func CreateField(column, row int) Field {
	return Field{
		column: column,
		row:    row,
	}
}

func (f Field) Column() int {
	return f.column
}

func (f Field) Row() int {
	return f.row
}

func (f Field) Left() Field {
	return CreateField(f.column-1, f.row)
}

func (f Field) Right() Field {
	return CreateField(f.column+1, f.row)
}

func (f Field) Up() Field {
	return CreateField(f.column, f.row-1)
}

func (f Field) Down() Field {
	return CreateField(f.column, f.row+1)
}
