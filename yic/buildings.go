package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
	"github.com/GodsBoss/ld46/pkg/grid/rect"
)

type building struct {
	typ       string
	gridXY    rect.Field
	x         float64
	y         float64
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
	building(p *playing, pos rect.Field) *building
}

type simplePlaceBuilding struct {
	cst         float64
	fldTypes    map[int]struct{}
	newBuilding func(p *playing, pos rect.Field) *building
}

func (pb *simplePlaceBuilding) cost() float64 {
	return pb.cst
}

func (pb *simplePlaceBuilding) fieldTypes() map[int]struct{} {
	return pb.fldTypes
}

func (pb *simplePlaceBuilding) building(p *playing, pos rect.Field) *building {
	return pb.newBuilding(p, pos)
}

var _ placeBuilding = &simplePlaceBuilding{}

var keyPlaceBuildingMapping = map[string]placeBuilding{
	"1": &simplePlaceBuilding{
		cst: 500.0,
		fldTypes: map[int]struct{}{
			fieldBuildSpot: struct{}{},
		},
		newBuilding: func(_ *playing, pos rect.Field) *building {
			return &building{
				typ:    "building_income",
				gridXY: pos,
				effect: &incomeBuildingEffect{},
			}
		},
	},
	"2": &simplePlaceBuilding{
		cst: 1000.0,
		fldTypes: map[int]struct{}{
			fieldBuildSpot: struct{}{},
		},
		newBuilding: func(p *playing, pos rect.Field) *building {
			x, y := p.levels.ChosenLevel().realCoordinateFloat64(float64(pos.Column()), float64(pos.Row()))
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
