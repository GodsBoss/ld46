package yic

import (
	"math/rand"

	"github.com/GodsBoss/ld46/pkg/engine"
)

type fx struct {
	x              int
	y              int
	key            string
	animation      float64
	animationSpeed float64
}

func (fx *fx) Tick(ms int) *engine.Transition {
	fx.animation += fx.animationSpeed * float64(ms) / 1000.0
	return nil
}

func (fx fx) over() bool {
	return fx.animation >= 1.0
}

type fxManager struct {
	currentID int
	fxs       map[int]*fx
}

func newFXManager() *fxManager {
	return &fxManager{
		currentID: 1,
		fxs:       make(map[int]*fx),
	}
}

func (m *fxManager) addFX(key string, x, y int) {
	newFX := &fx{
		x:              x,
		y:              y,
		key:            key,
		animationSpeed: animationSpeeds[key],
	}
	if newFX.animationSpeed == 0.0 {
		newFX.animationSpeed = defaultAnimationSpeed
	}
	m.fxs[m.currentID] = newFX
	m.currentID++
}

func (m *fxManager) addFXWithin(key string, left, top, right, bottom int) {
	m.addFX(key, rand.Intn(right-left)+left, rand.Intn(bottom-top)+top)
}

func (m *fxManager) Tick(ms int) *engine.Transition {
	fxOverList := make([]int, 0)
	for i := range m.fxs {
		m.fxs[i].Tick(ms)
		if m.fxs[i].over() {
			fxOverList = append(fxOverList, i)
		}
	}
	for i := range fxOverList {
		delete(m.fxs, fxOverList[i])
	}
	return nil
}

func (m *fxManager) Objects() []engine.Object {
	objects := make([]engine.Object, 0, len(m.fxs))
	for i := range m.fxs {
		objects = append(
			objects,
			engine.Object{
				Key:       m.fxs[i].key,
				X:         m.fxs[i].x,
				Y:         m.fxs[i].y,
				Animation: m.fxs[i].animation,
			},
		)
	}
	return objects
}

const (
	gunHit  = "gun_hit"
	gunShot = "gun_shot"
)

const defaultAnimationSpeed = 1.0

var animationSpeeds = map[string]float64{}
