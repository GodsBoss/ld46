// +build: js,wasm

package dom

import (
	"github.com/GodsBoss/ld46/pkg/errors"

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

func (canvas *Canvas) Context2D() (*Context2D, error) {
	jsCtx := canvas.value.Call("getContext", "2d")
	if jsCtx.IsNull() {
		return nil, errors.String("2d context not supported")
	}
	return &Context2D{
		value: jsCtx,
	}, nil
}

type Context2D struct {
	value js.Value
}

func (ctx2D *Context2D) DisableImageSmoothing() {
	ctx2D.value.Set("imageSmoothingEnabled", false)
}

func (ctx2D *Context2D) DrawImage(image *Image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight int) {
	ctx2D.value.Call("drawImage", image.value, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
}
