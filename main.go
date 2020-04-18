package main

import (
	"syscall/js"
	"time"

	"github.com/GodsBoss/ld46/pkg/console"
	"github.com/GodsBoss/ld46/pkg/dom"
	"github.com/GodsBoss/ld46/pkg/engine"
	"github.com/GodsBoss/ld46/pkg/engine/domevents"
	"github.com/GodsBoss/ld46/pkg/errors"
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
			// Run in background, else Browser main thread (for that window) will become unresponsive.
			go func() {
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
						game.ReceiveMouseEvent(domevents.FromMouseEvent(engine.MouseDown, event))
					},
				)
				dom.AddEventListener(
					canvas,
					"mouseup",
					func(event js.Value) {
						game.ReceiveMouseEvent(domevents.FromMouseEvent(engine.MouseUp, event))
					},
				)
				dom.AddEventListener(
					canvas,
					"mousemove",
					func(event js.Value) {
						game.ReceiveMouseEvent(domevents.FromMouseEvent(engine.MouseMove, event))
					},
				)
				ticker := time.NewTicker(time.Millisecond * 40)
				for {
					<-ticker.C
					game.Tick(40)
				}
			}()
		},
		func(err interface{}) {
			errsChan <- errors.String("loading game gfx failed")
			close(errsChan)
		},
	)
	return <-errsChan
}
