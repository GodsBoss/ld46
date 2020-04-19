package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type building struct {
	typ    string
	gridXY vector2D
	x      int
	y      int
	effect buildingEffect
}

func (b *building) Tick(ms int) *engine.Transition {
	return tickerTick(ms, b.effect)
}

type buildingEffect interface{}

type placeBuilding interface {
	cost() float64
	fieldTypes() map[int]struct{}
	building(p *playing, pos vector2D) *building
}

type simplePlaceBuilding struct {
	cst         float64
	fldTypes    map[int]struct{}
	newBuilding func(p *playing, pos vector2D) *building
}

func (pb *simplePlaceBuilding) cost() float64 {
	return pb.cst
}

func (pb *simplePlaceBuilding) fieldTypes() map[int]struct{} {
	return pb.fldTypes
}

func (pb *simplePlaceBuilding) building(p *playing, pos vector2D) *building {
	return pb.newBuilding(p, pos)
}

var _ placeBuilding = &simplePlaceBuilding{}

var keyPlaceBuildingMapping = map[string]placeBuilding{
	"1": &simplePlaceBuilding{
		cst: 600.0,
		fldTypes: map[int]struct{}{
			fieldBuildSpot: struct{}{},
		},
		newBuilding: func(_ *playing, pos vector2D) *building {
			return &building{
				typ:    "building_income",
				gridXY: pos,
				effect: &incomeBuildingEffect{},
			}
		},
	},
	"2": &simplePlaceBuilding{
		cst: 750.0,
		fldTypes: map[int]struct{}{
			fieldBuildSpot: struct{}{},
		},
		newBuilding: func(p *playing, pos vector2D) *building {
			b := &building{
				typ:    "building_gun",
				gridXY: pos,
				effect: &gun{
					p: p,
				},
			}
			return b
		},
	},
}
