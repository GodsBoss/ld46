package yic

import (
	"math"

	"github.com/GodsBoss/ld46/pkg/engine"
)

type head struct {
	p *playing

	phase     int
	animation float64
	health    float64
}

func (h *head) Init() {
	h.phase = phaseToddler
	h.animation = 0.0
	h.health = healthPerPhase[h.phase]
}

func (h *head) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	h.animation = math.Mod(h.animation+factor, 1.0)

	if h.health < 0.0 {
		if h.phase == 3 {
			return engine.NewTransition(gameOverStateID)
		}
		h.phase++
		h.health = healthPerPhase[h.phase]
		h.p.calculateIncomePerSecond()
	}

	return nil
}

func (h *head) IncomePerSecond() float64 {
	return baseResourcesPerSecondPerPhase[h.phase]
}

func (h *head) receiveDamage(dmg float64) {
	h.health -= dmg
}

func (h *head) key() string {
	return phaseHeadMapping[h.phase]
}

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

var healthPerPhase = map[int]float64{
	phaseToddler: 1000.0,
	phaseChild:   1000.0,
	phaseTeen:    1000.0,
}

var baseResourcesPerSecondPerPhase = map[int]float64{
	phaseToddler: 50.0,
	phaseChild:   40.0,
	phaseTeen:    30.0,
}
