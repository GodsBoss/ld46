package yic

import (
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
	resources       float64
	gridCursor      vector2D
}

func (p *playing) Init() {
	p.phase = 1
	p.resources = startResources
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
}

func (p *playing) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	p.headAnimation = math.Mod(p.headAnimation+factor, 1.0)
	p.resources += baseResourcesPerSecondPerPhase[p.phase] * factor
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

func (p *playing) HandleMouseEvent(event engine.MouseEvent) *engine.Transition {
	if event.Type == engine.MouseMove {
		p.gridCursor = p.levels.ChosenLevel().gridCursor(event.X, event.Y)
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
					X:   int(rx),
					Y:   int(ry),
				},
			)
		}
	}

	if p.levels.ChosenLevel().isOnGrid(p.gridCursor.X, p.gridCursor.Y) {
		cx, cy := lvl.realCoordinate(p.gridCursor.X, p.gridCursor.Y)
		objects["ui"] = append(
			objects["ui"],
			engine.Object{
				Key: "grid_cursor",
				X:   cx,
				Y:   cy,
			},
		)
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

const (
	phaseToddler = 1
	phaseChild   = 2
	phaseTeen    = 3
)

var phaseHeadMapping = map[int]string{
	phaseToddler: "head_toddler",
	phaseChild:   "head_child",
	phaseTeen:    "head_teen",
}

const healthPerPhase = 1000.0

const startResources = 1000.0

var baseResourcesPerSecondPerPhase = map[int]float64{
	phaseToddler: 100.0,
	phaseChild:   75.0,
	phaseTeen:    60.0,
}
