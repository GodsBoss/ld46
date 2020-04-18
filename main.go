package main

import (
	"github.com/GodsBoss/ld46/pkg/console"
	"github.com/GodsBoss/ld46/pkg/dom"
	"github.com/GodsBoss/ld46/pkg/errors"
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
			// TODO: Start game!
		},
		func(err interface{}) {
			errsChan <- errors.String("loading game gfx failed")
			close(errsChan)
		},
	)
	return <-errsChan
}
