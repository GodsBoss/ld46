package ui

import (
	"math"

	"github.com/GodsBoss/ld46/pkg/dom"
	"github.com/GodsBoss/ld46/pkg/engine"

	"sort"
)

type Renderer struct {
	GFXSource *dom.Image
	Sprites   map[string]Sprite
	Layers    []string
	Zoom      int
}

func NewRenderer(gfxSource *dom.Image, sprites map[string]Sprite, layers []string) *Renderer {
	return &Renderer{
		GFXSource: gfxSource,
		Sprites:   sprites,
		Layers:    layers,
	}
}

func (renderer *Renderer) zoom() int {
	z := renderer.Zoom
	if z < 1 {
		return 1
	}
	return z
}

func (renderer *Renderer) Draw(dest *dom.Context2D, objectsByLayer map[string][]engine.Object) {
	w, h := dest.Size()
	dest.ClearRect(0, 0, w, h)
	for i := range renderer.Layers {
		renderer.drawLayer(dest, objectsByLayer[renderer.Layers[i]])
	}
}

func (renderer *Renderer) drawLayer(dest *dom.Context2D, objects []engine.Object) {
	sortedObjects := ObjectsByZ(objects)
	sort.Sort(sortedObjects)
	for i := range sortedObjects {
		renderer.drawObject(dest, sortedObjects[i])
	}
}

func (renderer *Renderer) drawObject(dest *dom.Context2D, object engine.Object) {
	sprite := renderer.Sprites[object.Key]
	frame := 0
	if sprite.Frames > 1 {
		animation := math.Mod(object.Animation, 1.0)
		if animation < 0.0 {
			animation += 1.0
		}
		frame = int(animation * float64(sprite.Frames))
	}
	scale := 1
	if object.Scale > 1 {
		scale = object.Scale
	}
	dest.DrawImage(
		renderer.GFXSource,
		sprite.X+sprite.W*frame,
		sprite.Y,
		sprite.W,
		sprite.H,
		object.X*renderer.zoom(),
		object.Y*renderer.zoom(),
		sprite.W*renderer.zoom()*scale,
		sprite.H*renderer.zoom()*scale,
	)
}

// ObjectsByZ is a helper type for sorting objects by Z index.
type ObjectsByZ []engine.Object

func (objs ObjectsByZ) Len() int {
	return len(objs)
}

func (objs ObjectsByZ) Less(i, j int) bool {
	return objs[i].Z < objs[j].Z
}

func (objs ObjectsByZ) Swap(i, j int) {
	objs[i], objs[j] = objs[j], objs[i]
}

type Sprite struct {
	X int
	Y int

	W int
	H int

	Frames int
}
