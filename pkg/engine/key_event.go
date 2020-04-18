package engine

type KeyEvent struct {
	Type KeyEventType

	Alt   bool
	Ctrl  bool
	Shift bool

	Key string
}

type KeyEventType string

const (
	KeyUp    KeyEventType = "up"
	KeyDown  KeyEventType = "down"
	KeyPress KeyEventType = "press"
)

// KeyEventTarget can be implemented by States to receive key events.
type KeyEventTarget interface {
	HandleKeyEvent(event KeyEvent) *Transition
}
