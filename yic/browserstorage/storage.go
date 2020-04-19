package browserstorage

import (
	"github.com/GodsBoss/ld46/pkg/errors"

	"syscall/js"
)

type Storage struct{}

func (storage *Storage) Get(key string) (string, bool) {
	jsItem := js.Global().Get("localStorage").Call("getItem", key)
	if jsItem.IsNull() {
		return "", false
	}
	return jsItem.String(), true
}

func (storage *Storage) Set(key string, value string) error {
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				err = errors.String("could not write to localStorage")
			}
		}()
		js.Global().Get("localStorage").Call("setItem", key, value)
	}()
	return err
}

func (storage *Storage) Clear() {
	js.Global().Get("localStorage").Call("clear")
}
