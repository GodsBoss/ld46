package engine

type MouseEvent struct {
	Type MouseEventType

	Alt   bool
	Ctrl  bool
	Shift bool

	X int
	Y int

	// Button is the button pressed for mouse up and down events. Do not use this
	// directly, instead use PrimaryButton(), ...
	Button int
}

// PrimaryButton checks wether the primary button (usually left) was the cause
// of an up or down event.
func (event MouseEvent) PrimaryButton() bool {
	return (event.Type == MouseUp || event.Type == MouseDown) && event.Button == 0
}

// SecondaryButton checks wether the secondary button (usually right) was the
// cause of an up or down event.
func (event MouseEvent) SecondaryButton() bool {
	return (event.Type == MouseUp || event.Type == MouseDown) && event.Button == 2
}

// AuxiliaryButton checks wether the auxiliary button (usually middle / mouse wheel)
// was the cause of an up or down event.
func (event MouseEvent) AuxiliaryButton() bool {
	return (event.Type == MouseUp || event.Type == MouseDown) && event.Button == 1
}

type MouseEventType string

const (
	MouseUp   MouseEventType = "up"
	MouseDown MouseEventType = "down"
	MouseMove MouseEventType = "move"
)

// MouseEventTarget can be implemented by States to receive mouse events.
type MouseEventTarget interface {
	HandleMouseEvent(event MouseEvent) *Transition
}
