package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels *levels

	head *head

	responsibilites map[int][]*responsibility
	phase           int
	resources       float64
	incomePerSecond float64
	gridCursor      vector2D
	buildings       map[vector2D]*building
}

func (p *playing) Init() {
	p.head = &head{
		p: p,
	}
	p.head.Init()
	p.buildings = make(map[vector2D]*building)
	p.resources = startResources
	p.calculateIncomePerSecond()
	p.responsibilites = make(map[int][]*responsibility)
	for chainIndex := range p.levels.ChosenLevel().chains {
		p.responsibilites[chainIndex] = make([]*responsibility, 0)
	}
	p.responsibilites[0] = append(
		p.responsibilites[0],
		&responsibility{
			typ:      responsibilityType1,
			speed:    12.5,
			position: -5.0,
			life:     1500,
		},
	)
}

func (p *playing) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	p.resources += p.incomePerSecond * factor
	for chainIndex := range p.responsibilites {
		respsToRemove := make(map[int]struct{})
		for i := range p.responsibilites[chainIndex] {
			var headReached bool
			resp := p.responsibilites[chainIndex][i]
			resp.position += resp.speed * factor
			resp.x, resp.y, headReached = p.levels.ChosenLevel().responsibilityPosition(chainIndex, resp.position)
			if headReached {
				respsToRemove[i] = struct{}{}
			}
		}
		if len(respsToRemove) > 0 {
			remaining := make([]*responsibility, 0, len(p.responsibilites)-len(respsToRemove))
			for i := range p.responsibilites[chainIndex] {
				resp := p.responsibilites[chainIndex][i]
				if _, okRemove := respsToRemove[i]; okRemove {
					p.head.receiveDamage(resp.life)
				} else {
					remaining = append(remaining, resp)
				}
			}
			p.responsibilites[chainIndex] = remaining
		}
	}
	return p.head.Tick(ms)
}

func (p *playing) calculateIncomePerSecond() {
	p.incomePerSecond = p.head.IncomePerSecond()
	for v := range p.buildings {
		if provider, ok := p.buildings[v].effect.(incomeProviderBuilding); ok {
			p.incomePerSecond += provider.IncomePerSecond()
		}
	}
}

type incomeProviderBuilding interface {
	IncomePerSecond() float64
}

func (p *playing) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	if event.Key == "x" {
		return engine.NewTransition(gameOverStateID)
	}
	if placeBuilding, ok := keyPlaceBuildingMapping[event.Key]; ok {
		// Too expensive.
		if placeBuilding.cost() > p.resources {
			return nil
		}

		lvl := p.levels.ChosenLevel()

		// Far from any field.
		if !lvl.isOnGrid(p.gridCursor.X, p.gridCursor.Y) {
			return nil
		}
		// Field type not matching.
		if _, ok := placeBuilding.fieldTypes()[lvl.fields[p.gridCursor.Y][p.gridCursor.X].typ]; !ok {
			return nil
		}
		// Field already contains building.
		if _, ok := p.buildings[p.gridCursor]; ok {
			return nil
		}

		// Finally, build building.
		p.buildings[p.gridCursor] = placeBuilding.building(p.gridCursor)
		x, y := p.levels.ChosenLevel().realCoordinateFloat64(float64(p.gridCursor.X), float64(p.gridCursor.Y))
		p.buildings[p.gridCursor].x = int(x)
		p.buildings[p.gridCursor].y = int(y)
		p.resources -= placeBuilding.cost()

		p.calculateIncomePerSecond()

		return nil
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
				Key:       p.head.key(),
				X:         headX,
				Y:         headY,
				Animation: p.head.animation,
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

	for v := range p.buildings {
		objects["entities"] = append(
			objects["entities"],
			engine.Object{
				Key: p.buildings[v].typ,
				X:   p.buildings[v].x,
				Y:   p.buildings[v].y,
			},
		)
	}

	for chainIndex := range p.responsibilites {
		for i := range p.responsibilites[chainIndex] {
			objects["entities"] = append(
				objects["entities"],
				engine.Object{
					Key: p.responsibilites[chainIndex][i].typ,
					X:   int(p.responsibilites[chainIndex][i].x),
					Y:   int(p.responsibilites[chainIndex][i].y),
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

	// x and y are calculated via position.
	x float64
	y float64
}

const (
	responsibilityType1 = "responsibility_1"
	responsibilityType2 = "responsibility_2"
	responsibilityType3 = "responsibility_3"
)

const startResources = 1000.0
