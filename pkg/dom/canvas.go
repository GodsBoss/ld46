// +build: js,wasm

package dom

import (
	"syscall/js"
)

// Canvas wraps a JS canvas element. Canvas implements Node.
type Canvas struct {
	value js.Value
}

func (canvas *Canvas) getJSNode() js.Value {
	return canvas.value
}

func (canvas *Canvas) SetSize(width, height int) {
	canvas.SetWidth(width)
	canvas.SetHeight(height)
}

func (canvas *Canvas) SetWidth(width int) {
	canvas.value.Set("width", width)
}

func (canvas *Canvas) SetHeight(height int) {
	canvas.value.Set("height", height)
}
