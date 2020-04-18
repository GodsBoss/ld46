package main

import (
	"syscall/js"
	"time"

	"github.com/GodsBoss/ld46/pkg/console"
	"github.com/GodsBoss/ld46/pkg/dom"
	"github.com/GodsBoss/ld46/pkg/engine"
	"github.com/GodsBoss/ld46/pkg/engine/domevents"
	"github.com/GodsBoss/ld46/pkg/errors"
	"github.com/GodsBoss/ld46/pkg/ui"
	"github.com/GodsBoss/ld46/yic"
)

func main() {
	err := run()
	if err != nil {
		console.Global().LogMessage(err.Error())
	}
}

func run() error {
	w, err := dom.GlobalWindow()
	if err != nil {
		return err
	}
	doc, err := w.Document()
	if err != nil {
		return err
	}
	canvas, err := doc.CreateCanvasElement()
	if err != nil {
		return err
	}
	canvas.SetSize(800, 600)
	gameElement, err := doc.GetElementByID("game")
	if err != nil {
		return err
	}
	err = gameElement.AppendChild(canvas)
	if err != nil {
		return err
	}
	img, err := doc.CreateImageElement("gfx.png")
	if err != nil {
		return err
	}
	errsChan := make(chan error)
	img.On(
		func() {
			game := yic.NewGame()
			dom.AddEventListener(
				w,
				"keydown",
				func(event js.Value) {
					game.ReceiveKeyEvent(domevents.FromKeyEvent(engine.KeyDown, event))
				},
			)
			dom.AddEventListener(
				w,
				"keyup",
				func(event js.Value) {
					game.ReceiveKeyEvent(domevents.FromKeyEvent(engine.KeyUp, event))
				},
			)
			dom.AddEventListener(
				canvas,
				"mousedown",
				func(event js.Value) {
					game.ReceiveMouseEvent(translateMouseEvent(zoom, domevents.FromMouseEvent(engine.MouseDown, event)))
				},
			)
			dom.AddEventListener(
				canvas,
				"mouseup",
				func(event js.Value) {
					game.ReceiveMouseEvent(translateMouseEvent(zoom, domevents.FromMouseEvent(engine.MouseUp, event)))
				},
			)
			dom.AddEventListener(
				canvas,
				"mousemove",
				func(event js.Value) {
					game.ReceiveMouseEvent(translateMouseEvent(zoom, domevents.FromMouseEvent(engine.MouseMove, event)))
				},
			)

			cancelGameLoop := make(chan struct{})
			// Run in background, else Browser main thread (for that window) will become unresponsive.
			go func() {
				ticker := time.NewTicker(time.Millisecond * msPerTick)
				for {
					select {
					case <-ticker.C:
					case <-cancelGameLoop:
						return
					}
					game.Tick(msPerTick)
				}
			}()

			// Drawing gets its own loop via requestAnimationFrame.
			context2D, err := canvas.Context2D()
			if err != nil {
				close(cancelGameLoop)
				errsChan <- err
				close(errsChan)
			}
			renderer := ui.NewRenderer(img, yic.Sprites(), []string{"background", "fields", "entities", "fx", "ui"})
			renderer.Zoom = zoom
			var reqAnimationFrameCallback func()
			reqAnimationFrameCallback = func() {
				w.RequestAnimationFrame(reqAnimationFrameCallback)
				renderer.Draw(context2D, game.Objects())
			}
			w.RequestAnimationFrame(reqAnimationFrameCallback)
		},
		func(err interface{}) {
			errsChan <- errors.String("loading game gfx failed")
			close(errsChan)
		},
	)
	return <-errsChan
}

const msPerTick = 40

const zoom = 2

func translateMouseEvent(zoom int, event engine.MouseEvent) engine.MouseEvent {
	event.X = event.X / zoom
	event.Y = event.Y / zoom
	return event
}
