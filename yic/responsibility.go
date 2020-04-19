package yic

import (
	"math/rand"

	"github.com/GodsBoss/ld46/pkg/engine"
)

type responsibilities struct {
	p *playing

	byChain map[int][]*responsibility

	spawnBuffer        float64
	spawnSpeed         float64
	spawnSpeedIncrease float64
	spawnLife          float64
	spawnType          string
	spawnReward        float64
}

func (resps *responsibilities) Init() {
	resps.byChain = make(map[int][]*responsibility)
	for chainIndex := range resps.p.levels.ChosenLevel().chains {
		resps.byChain[chainIndex] = make([]*responsibility, 0)
	}

	resps.spawnSpeedIncrease = 0.05
	resps.spawnSpeed = initialSpawnSpeed
	resps.spawnLife = initialSpawnLife
	resps.spawnType = responsibilityType1
	resps.spawnReward = initialSpawnReward
}

func (resps *responsibilities) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0

	for chainIndex := range resps.byChain {
		for i := range resps.byChain[chainIndex] {
			resps.byChain[chainIndex][i].animation += factor
		}
	}

	// Check for resps without health and remove them.
	for chainIndex := range resps.byChain {
		respsWithoutHealth := make(map[int]struct{})
		for i := range resps.byChain[chainIndex] {
			if resps.byChain[chainIndex][i].life <= 0 {
				respsWithoutHealth[i] = struct{}{}
			}
		}
		if len(respsWithoutHealth) > 0 {
			remaining := make([]*responsibility, 0)
			for i := range resps.byChain[chainIndex] {
				if _, ok := respsWithoutHealth[i]; !ok {
					remaining = append(remaining, resps.byChain[chainIndex][i])
				} else {
					resps.p.resources += resps.byChain[chainIndex][i].reward
				}
			}
			resps.byChain[chainIndex] = remaining
		}
	}

	// Move all resps, then check wether they reached head.
	for chainIndex := range resps.byChain {
		respsAtHead := make(map[int]struct{})
		for i := range resps.byChain[chainIndex] {
			var headReached bool
			resp := resps.byChain[chainIndex][i]
			resp.position += resp.speed * factor
			resp.x, resp.y, headReached = resps.p.levels.ChosenLevel().responsibilityPosition(chainIndex, resp.position)
			if headReached {
				respsAtHead[i] = struct{}{}
			}
		}
		if len(respsAtHead) > 0 {
			remaining := make([]*responsibility, 0, len(resps.byChain)-len(respsAtHead))
			for i := range resps.byChain[chainIndex] {
				resp := resps.byChain[chainIndex][i]
				if _, okRemove := respsAtHead[i]; okRemove {
					resps.p.head.receiveDamage(resp.life)
				} else {
					remaining = append(remaining, resp)
				}
			}
			resps.byChain[chainIndex] = remaining
		}
	}

	// Spawn new resps.
	resps.spawnSpeed += resps.spawnSpeedIncrease * factor
	if resps.spawnSpeed > spawnSpeedThreshold {
		resps.spawnSpeed = initialSpawnSpeed
		resps.spawnLife += spawnLifeIncrease
		resps.spawnType = nextSpawnType[resps.spawnType]
		resps.spawnReward += spawnRewardIncrease
	}
	resps.spawnBuffer += resps.spawnSpeed * factor
	if resps.spawnBuffer > 1.0 {
		chainIndex := rand.Intn(len(resps.byChain))
		resps.byChain[chainIndex] = append(
			resps.byChain[chainIndex],
			&responsibility{
				typ:    resps.spawnType,
				life:   resps.spawnLife,
				speed:  0.5 + 0.5*rand.Float64(),
				reward: resps.spawnReward,
			},
		)
		resps.spawnBuffer -= 1.0
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
					Key:       resps.byChain[chainIndex][i].typ,
					X:         int(resps.byChain[chainIndex][i].x),
					Y:         int(resps.byChain[chainIndex][i].y),
					Z:         int(resps.byChain[chainIndex][i].y * 1000.0),
					Animation: resps.byChain[chainIndex][i].animation,
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

	// reward is added to the player's resources when this responsibility is killed.
	reward float64

	animation float64
}

func (r *responsibility) receiveDamage(dmg float64) {
	r.life -= dmg
}

const (
	responsibilityType1 = "responsibility_1"
	responsibilityType2 = "responsibility_2"
	responsibilityType3 = "responsibility_3"
)

var nextSpawnType = map[string]string{
	responsibilityType1: responsibilityType2,
	responsibilityType2: responsibilityType3,
	responsibilityType3: responsibilityType1,
}

const initialSpawnSpeed = 1.0
const initialSpawnLife = 250.0
const initialSpawnReward = 25.0

// spawnSpeedThreshold is the threshold for the spawn speed to reset, increase spawn's life and rewards, and switch the spawn type.
const spawnSpeedThreshold = 1.5

const spawnLifeIncrease = 50.0
const spawnRewardIncrease = 5.0
