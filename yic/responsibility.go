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
	spawnType          string
	enemiesKilled      int
	wave               int

	defaultTemplate *responsibilityTemplate
}

func (resps *responsibilities) Init() {
	resps.byChain = make(map[int][]*responsibility)
	for chainIndex := range resps.p.levels.ChosenLevel().chains {
		resps.byChain[chainIndex] = make([]*responsibility, 0)
	}

	resps.spawnSpeedIncrease = 0.05
	resps.spawnSpeed = initialSpawnSpeed
	resps.spawnType = responsibilityType1

	resps.defaultTemplate = &responsibilityTemplate{
		baseLife:    250.0,
		lifePerWave: 50.0,

		baseSpeed:           0.5,
		potentialSpeedBoost: 0.5,

		baseReward:    25.0,
		rewardPerWave: 5.0,
	}

	resps.enemiesKilled = 0
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
					resps.enemiesKilled++
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
		resps.wave++
		resps.spawnSpeed = initialSpawnSpeed
		resps.spawnType = nextSpawnType[resps.spawnType]
	}
	resps.spawnBuffer += resps.spawnSpeed * factor
	if resps.spawnBuffer > 1.0 {
		chainIndex := rand.Intn(len(resps.byChain))
		x, y, _ := resps.p.levels.ChosenLevel().responsibilityPosition(chainIndex, 0)
		resp := resps.defaultTemplate.responsibilityByWave(resps.wave, x, y)
		resp.typ = resps.spawnType
		resps.byChain[chainIndex] = append(resps.byChain[chainIndex], resp)
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

type responsibilityTemplate struct {
	baseLife    float64
	lifePerWave float64

	baseSpeed           float64
	potentialSpeedBoost float64

	baseReward    float64
	rewardPerWave float64
}

func (tmpl *responsibilityTemplate) responsibilityByWave(wave int, x, y float64) *responsibility {
	return &responsibility{
		life:   tmpl.baseLife + (float64(wave))*tmpl.lifePerWave,
		speed:  tmpl.baseSpeed + tmpl.potentialSpeedBoost*rand.Float64(),
		reward: tmpl.baseReward + (float64(wave))*tmpl.rewardPerWave,
		x:      x,
		y:      y,
	}
}

const (
	responsibilityType1 = "responsibility_1"
	responsibilityType2 = "responsibility_2"
	responsibilityType3 = "responsibility_3"
	responsibilityType4 = "responsibility_4"
)

var nextSpawnType = map[string]string{
	responsibilityType1: responsibilityType2,
	responsibilityType2: responsibilityType3,
	responsibilityType3: responsibilityType4,
	responsibilityType4: responsibilityType1,
}

const initialSpawnSpeed = 1.0
const initialSpawnLife = 250.0
const initialSpawnReward = 25.0

// spawnSpeedThreshold is the threshold for the spawn speed to reset, increase spawn's life and rewards, and switch the spawn type.
const spawnSpeedThreshold = 1.5

const spawnLifeIncrease = 50.0
const spawnRewardIncrease = 5.0
