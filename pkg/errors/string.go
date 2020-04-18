package errors

// String is a simple string error.
type String string

func (s String) Error() string {
	return string(s)
}
