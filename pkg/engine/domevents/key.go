// +build: js,wasm

package domevents

import (
	"github.com/GodsBoss/ld46/pkg/engine"

	"syscall/js"
)

func FromKeyEvent(typ engine.KeyEventType, domEvent js.Value) engine.KeyEvent {
	return engine.KeyEvent{
		Type:  typ,
		Alt:   domEvent.Get("altKey").Bool(),
		Ctrl:  domEvent.Get("ctrlKey").Bool(),
		Shift: domEvent.Get("shiftKey").Bool(),
		Key:   domEvent.Get("key").String(),
	}
}
