package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type ticker interface {
	Tick(ms int) *engine.Transition
}

func tickerTick(ms int, candidate interface{}) *engine.Transition {
	if t, ok := candidate.(ticker); ok {
		return t.Tick(ms)
	}
	return nil
}
