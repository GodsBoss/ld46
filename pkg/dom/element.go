// +build: js,wasm

package dom

import (
	"syscall/js"
)

type Element struct {
	value js.Value
}

func (element *Element) AppendChild(node Node) error {
	element.value.Call("appendChild", node.getJSNode())
	return nil
}

type Node interface {
	getJSNode() js.Value
}
