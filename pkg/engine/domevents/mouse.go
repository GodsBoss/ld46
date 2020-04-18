// +build: js,wasm

package domevents

import (
	"github.com/GodsBoss/ld46/pkg/engine"

	"syscall/js"
)

func FromMouseEvent(typ engine.MouseEventType, domEvent js.Value) engine.MouseEvent {
	target := domEvent.Get("target")
	return engine.MouseEvent{
		Type:   typ,
		Alt:    domEvent.Get("altKey").Bool(),
		Ctrl:   domEvent.Get("ctrlKey").Bool(),
		Shift:  domEvent.Get("shiftKey").Bool(),
		X:      domEvent.Get("clientX").Int() - target.Get("offsetLeft").Int(),
		Y:      domEvent.Get("clientY").Int() - target.Get("offsetTop").Int(),
		Button: domEvent.Get("button").Int(),
	}
}
