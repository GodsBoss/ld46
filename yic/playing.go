package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels *levels

	head *head

	fxManager   *fxManager
	textManager *textManager

	responsibilites *responsibilities
	phase           int
	resources       float64
	incomePerSecond float64
	gridCursor      vector2D
	buildings       map[vector2D]*building
}

func (p *playing) Init() {
	p.fxManager = newFXManager()
	p.textManager = newTextManager()
	p.head = &head{
		p: p,
	}
	p.head.Init()
	p.buildings = make(map[vector2D]*building)
	p.resources = startResources
	p.calculateIncomePerSecond()
	p.responsibilites = &responsibilities{
		p: p,
	}
	p.responsibilites.Init()
}

func (p *playing) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	p.resources += p.incomePerSecond * factor
	p.responsibilites.Tick(ms)
	for v := range p.buildings {
		p.buildings[v].Tick(ms)
	}
	p.fxManager.Tick(ms)
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
		p.buildings[p.gridCursor] = placeBuilding.building(p, p.gridCursor)
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
		"fx": p.fxManager.Objects(),
		"ui": p.textManager.Objects(),
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
				Key:       p.buildings[v].typ,
				X:         p.buildings[v].x,
				Y:         p.buildings[v].y,
				Animation: p.buildings[v].animation,
			},
		)
	}

	objects["entities"] = append(objects["entities"], p.responsibilites.Objects()...)

	if p.levels.ChosenLevel().isOnGrid(p.gridCursor.X, p.gridCursor.Y) {
		cx, cy := lvl.realCoordinate(p.gridCursor.X, p.gridCursor.Y)
		key := "grid_cursor_no"
		if p.levels.ChosenLevel().fields[p.gridCursor.Y][p.gridCursor.X].typ == fieldBuildSpot {
			key = "grid_cursor"
			if p.buildings[p.gridCursor] == nil {
				key = "grid_cursor_yes"
			}
		}
		objects["ui"] = append(
			objects["ui"],
			engine.Object{
				Key: key,
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

const startResources = 1000.0
