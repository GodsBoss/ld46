// +build: js,wasm

package dom

import (
	"syscall/js"
)

func AddEventListener(node Node, typ string, listener func(event js.Value)) {
	node.getJSNode().Call(
		"addEventListener",
		typ,
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				listener(args[0])
				return nil
			},
		),
		false,
	)
}
