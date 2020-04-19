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

func (element *Element) getJSNode() js.Value {
	return element.value
}

type Node interface {
	getJSNode() js.Value
}

func RemoveNode(node Node) {
	node.getJSNode().Get("parentNode").Call("removeChild", node.getJSNode())
}
