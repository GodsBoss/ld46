// +build: js,wasm

package dom

import (
	"syscall/js"
)

type Image struct {
	value js.Value
}

func (image *Image) getJSNode() js.Value {
	return image.value
}

func (image *Image) On(load func(), fail func(err interface{})) {
	image.value.Call(
		"addEventListener",
		"load",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				load()
				return nil
			},
		),
		false,
	)

	image.value.Call(
		"addEventListener",
		"error",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				var err interface{}
				if len(args) > 0 {
					err = args[0]
				}
				fail(err)
				return nil
			},
		),
		false,
	)
}
