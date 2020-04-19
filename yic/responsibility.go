package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type responsibilities struct {
	p *playing

	byChain map[int][]*responsibility
}

func (resps *responsibilities) Init() {
	resps.byChain = make(map[int][]*responsibility)
	for chainIndex := range resps.p.levels.ChosenLevel().chains {
		resps.byChain[chainIndex] = make([]*responsibility, 0)
	}
	resps.byChain[0] = append(
		resps.byChain[0],
		&responsibility{
			typ:   responsibilityType1,
			life:  1500,
			speed: 2.5,
		},
	)
}

func (resps *responsibilities) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	for chainIndex := range resps.byChain {
		respsToRemove := make(map[int]struct{})
		for i := range resps.byChain[chainIndex] {
			var headReached bool
			resp := resps.byChain[chainIndex][i]
			resp.position += resp.speed * factor
			resp.x, resp.y, headReached = resps.p.levels.ChosenLevel().responsibilityPosition(chainIndex, resp.position)
			if headReached {
				respsToRemove[i] = struct{}{}
			}
		}
		if len(respsToRemove) > 0 {
			remaining := make([]*responsibility, 0, len(resps.byChain)-len(respsToRemove))
			for i := range resps.byChain[chainIndex] {
				resp := resps.byChain[chainIndex][i]
				if _, okRemove := respsToRemove[i]; okRemove {
					resps.p.head.receiveDamage(resp.life)
				} else {
					remaining = append(remaining, resp)
				}
			}
			resps.byChain[chainIndex] = remaining
		}
	}
	return nil
}

func (resps *responsibilities) Objects() []engine.Object {
	objects := make([]engine.Object, 0)
	for chainIndex := range resps.byChain {
		for i := range resps.byChain[chainIndex] {
			objects = append(
				objects,
				engine.Object{
					Key: resps.byChain[chainIndex][i].typ,
					X:   int(resps.byChain[chainIndex][i].x),
					Y:   int(resps.byChain[chainIndex][i].y),
				},
			)
		}
	}
	return objects
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
