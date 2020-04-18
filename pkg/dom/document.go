// +build: js,wasm

package dom

import (
	"github.com/GodsBoss/ld46/pkg/errors"

	"syscall/js"
)

type Document struct {
	value js.Value
}

func (doc *Document) GetElementByID(id string) (*Element, error) {
	jsEl := doc.value.Call("getElementById", id)
	if jsEl.IsNull() {
		return nil, errors.String("element with id " + id + " does not exist")
	}
	return &Element{
		value: jsEl,
	}, nil
}

func (doc *Document) CreateCanvasElement() (*Canvas, error) {
	jsCanvas := doc.value.Call("createElement", "canvas")
	if jsCanvas.IsNull() {
		return nil, errors.String("could not create canvas element")
	}
	return &Canvas{
		value: jsCanvas,
	}, nil
}
