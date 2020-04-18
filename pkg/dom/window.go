// +build: js,wasm

package dom

import (
	"github.com/GodsBoss/ld46/pkg/errors"

	"syscall/js"
)

type Window struct {
	value js.Value
}

func GlobalWindow() (*Window, error) {
	glob := js.Global()
	if glob.IsNull() {
		return nil, errors.String("window object does not exist")
	}
	return &Window{
		value: glob,
	}, nil
}

func (w *Window) getValue() js.Value {
	v := w.value
	if !v.IsNull() {
		return v
	}
	return js.Global()
}

func (w *Window) getJSNode() js.Value {
	return w.value
}

func (w *Window) Document() (*Document, error) {
	jsDoc := w.getValue().Get("document")
	if jsDoc.IsNull() {
		return nil, errors.String("document object does not exist")
	}
	return &Document{
		value: jsDoc,
	}, nil
}

func (w *Window) RequestAnimationFrame(f func()) {
	w.getValue().Call(
		"requestAnimationFrame",
		js.FuncOf(
			func(_ js.Value, _ []js.Value) interface{} {
				f()
				return nil
			},
		),
	)
}
