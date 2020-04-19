package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type gun struct {
	p *playing

	reloading float64
}

func (g *gun) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0
	g.reloading -= factor
	if g.reloading < 0 {
		g.reloading = 0
	}

	// Gun is still reloading and therefore, cannot shoot.
	if g.reloading > 0 {
		return nil
	}

	return nil
}

const gunRange = 100.0
