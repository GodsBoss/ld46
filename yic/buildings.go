package yic

type building interface {
	gridXY() vector2D
}

type placeBuilding interface {
	cost() float64
	fieldTypes() map[int]struct{}
	building(pos vector2D) building
}

type simplePlaceBuilding struct {
	cst         float64
	fldTypes    map[int]struct{}
	newBuilding func(pos vector2D) building
}

func (pb *simplePlaceBuilding) cost() float64 {
	return pb.cst
}

func (pb *simplePlaceBuilding) fieldTypes() map[int]struct{} {
	return pb.fldTypes
}

func (pb *simplePlaceBuilding) building(pos vector2D) building {
	return pb.newBuilding(pos)
}

var _ placeBuilding = &simplePlaceBuilding{}

var keyPlaceBuildingMapping = map[string]placeBuilding{
	"1": &simplePlaceBuilding{
		cst: 600.0,
		fldTypes: map[int]struct{}{
			fieldBuildSpot: struct{}{},
		},
		newBuilding: func(pos vector2D) building {
			return &incomeBuilding{
				pos: pos,
			}
		},
	},
}
