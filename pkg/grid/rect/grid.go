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

// CoordinatesFromField takes a field and returns the corresponding coordinates of the top left corner
// as a vector. Unlike FieldFromCoordinates, CoordinatesFromField works on grids with zero field size,
// it is just not useful because all grid points have the same coordinates: (0, 0).
func (grid *Grid) CoordinatesFromField(field Field) v2d.Vector {
	return v2d.FromXY(
		grid.fieldSize.X()*float64(field.Column())+grid.offset.X(),
		grid.fieldSize.Y()*float64(field.Row())+grid.offset.Y(),
	)
}

// Field is a single field of a grid, identified by column and row. Fields are value objects, i.e. they never change.
// Fields can safely be used as map keys.
type Field struct {
	column int
	row    int
}

// CreateField creates a field point to (column, row).
func CreateField(column, row int) Field {
	return Field{
		column: column,
		row:    row,
	}
}

// Column returns the field's column.
func (f Field) Column() int {
	return f.column
}

// Row returns the field's row.
func (f Field) Row() int {
	return f.row
}

// FieldOffset is an offset for fields, e.g. a linear translation.
type FieldOffset interface {
	// Apply applies this offset to a field, resulting in a new field. The old field
	// remains unchanged.
	Apply(Field) Field
}

// FieldOffsetFromField converts a field to a corresponding offset, meaning its
// column is used as the horizontal part of the offset, while the row is used as
// the vertical part.
func FieldOffsetFromField(f Field) FieldOffset {
	return fieldFieldOffset(f)
}

type fieldFieldOffset Field

func (offset fieldFieldOffset) Apply(f Field) Field {
	f.column += offset.column
	f.row += offset.row
	return f
}

var _ FieldOffset = fieldFieldOffset{}

// FieldOffsetLeft returns an offset which takes the field left to the given field.
func FieldOffsetLeft() FieldOffset {
	return FieldOffsetFromField(CreateField(-1, 0))
}

// FieldOffsetRight returns an offset which takes the field right to the given field.
func FieldOffsetRight() FieldOffset {
	return FieldOffsetFromField(CreateField(1, 0))
}

// FieldOffsetUp returns an offset which takes the field above the given field.
func FieldOffsetUp() FieldOffset {
	return FieldOffsetFromField(CreateField(0, -1))
}

// FieldOffsetDown returns an offset which takes the field below the given field.
func FieldOffsetDown() FieldOffset {
	return FieldOffsetFromField(CreateField(0, 1))
}
