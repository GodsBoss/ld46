package yic

import (
	"fmt"
	"math"

	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels *levels

	headAnimation   float64
	responsibilites map[int][]*responsibility
	phase           int
	headHealth      float64
}

func (p *playing) Init() {
	p.phase = 1
	p.headAnimation = 0.0
	p.headHealth = healthPerPhase
	p.responsibilites = make(map[int][]*responsibility)
	for chainIndex := range p.levels.ChosenLevel().chains {
		p.responsibilites[chainIndex] = make([]*responsibility, 0)
	}
	p.responsibilites[0] = append(
		p.responsibilites[0],
		&responsibility{
			typ:      responsibilityType1,
			speed:    2.5,
			position: -5.0,
		},
	)
	fmt.Println(p.levels.ChosenLevel().responsibilityPosition(0, 0))
}

func (p *playing) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	p.headAnimation = math.Mod(p.headAnimation+factor, 1.0)
	for chainIndex := range p.responsibilites {
		for i := range p.responsibilites[chainIndex] {
			p.responsibilites[chainIndex][i].position += p.responsibilites[chainIndex][i].speed * factor
		}
	}
	if p.headHealth < 0.0 {
		if p.phase == 3 {
			return engine.NewTransition(gameOverStateID)
		}
		p.phase++
		p.headHealth = healthPerPhase
	}
	return nil
}

func (p *playing) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type == engine.KeyUp && event.Key == "x" {
		return engine.NewTransition(gameOverStateID)
	}
	return nil
}

func (p *playing) Objects() map[string][]engine.Object {
	lvl := p.levels.ChosenLevel()
	headX, headY := lvl.realCoordinate(lvl.headX, lvl.headY)
	objects := map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_playing",
				X:   0,
				Y:   0,
			},
		},
		"fields": []engine.Object{},
		"entities": []engine.Object{
			engine.Object{
				Key:       phaseHeadMapping[p.phase],
				X:         headX,
				Y:         headY,
				Animation: p.headAnimation,
			},
		},
	}

	for row := range lvl.fields {
		for col := range lvl.fields[row] {
			field := lvl.fields[row][col]
			key := fieldTypeSpriteKeyMapping[field.typ]
			if key == "" {
				continue
			}
			rx, ry := lvl.realCoordinate(col, row)
			objects["fields"] = append(
				objects["fields"],
				engine.Object{
					Key: key,
					X:   rx,
					Y:   ry,
				},
			)
		}
	}

	for chainIndex := range p.responsibilites {
		for i := range p.responsibilites[chainIndex] {
			rx, ry := p.levels.ChosenLevel().responsibilityPosition(chainIndex, p.responsibilites[chainIndex][i].position)
			objects["entities"] = append(
				objects["entities"],
				engine.Object{
					Key: p.responsibilites[chainIndex][i].typ,
					X:   rx,
					Y:   ry,
				},
			)
		}
	}

	return objects
}

var fieldTypeSpriteKeyMapping = map[int]string{
	fieldWay:       "field_way",
	fieldObstacle:  "field_obstacle",
	fieldBuildSpot: "field_buildspot",
}

var fieldSize = vector2D{
	X: 18,
	Y: 18,
}

type responsibility struct {
	typ   string
	life  float64
	speed float64

	// position is the position of the responsibility on its chain.
	position float64
}

const (
	responsibilityType1 = "responsibility_1"
	responsibilityType2 = "responsibility_2"
	responsibilityType3 = "responsibility_3"
)

var phaseHeadMapping = map[int]string{
	1: "head_toddler",
	2: "head_child",
	3: "head_teen",
}

const healthPerPhase = 1000.0
