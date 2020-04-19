package yic

import (
	"math"

	"github.com/GodsBoss/ld46/pkg/engine"
)

type gun struct {
	p *playing

	x float64
	y float64

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

	for _, chain := range g.p.responsibilites.byChain {
		for _, resp := range chain {
			dist := math.Sqrt((g.x-resp.x)*(g.x-resp.x) + (g.y-resp.y)*(g.y-resp.y))
			if dist <= gunRange {
				// Shoot!
				g.reloading = gunReload
				resp.receiveDamage(gunDmg)
				g.p.fxManager.addFXWithin("gun_shot", int(g.x), int(g.y), int(g.x)+14, int(g.y)+14)
				g.p.fxManager.addFXWithin("gun_hit", int(resp.x), int(resp.y), int(resp.x)+12, int(resp.y)+12)
			}
		}
	}

	return nil
}

const gunRange = 40.0
const gunDmg = 75.0

// gunReload is the gun's reloading time in seconds.
const gunReload = 0.5
