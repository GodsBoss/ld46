package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type building struct {
	typ       string
	gridXY    vector2D
	x         int
	y         int
	effect    buildingEffect
	animation float64
}

func (b *building) Tick(ms int) *engine.Transition {
	b.animation += float64(ms) / 1000.0
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
			x, y := p.levels.ChosenLevel().realCoordinateFloat64(float64(pos.X), float64(pos.Y))
			b := &building{
				typ:    "building_gun",
				gridXY: pos,
				effect: &gun{
					p: p,
					x: x,
					y: y,
				},
			}
			return b
		},
	},
}
