package main

import (
	"github.com/GodsBoss/ld46/pkg/console"
	"github.com/GodsBoss/ld46/pkg/dom"
)

func main() {
	err := run()
	if err != nil {
		console.Global().LogObjects(err)
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
	return gameElement.AppendChild(canvas)
}
