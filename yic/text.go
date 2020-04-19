package yic

import (
	"strings"

	"github.com/GodsBoss/ld46/pkg/engine"
)

type textManager struct {
	texts     map[string]*text
	stringMap func(string) string
}

func newTextManager() *textManager {
	filter := createFilter(allowedChars)
	sm := func(input string) string {
		return strings.Map(
			filter,
			strings.ToLower(input),
		)
	}
	return &textManager{
		texts:     make(map[string]*text),
		stringMap: sm,
	}
}

func (m *textManager) Objects() []engine.Object {
	objects := make([]engine.Object, 0)
	for id := range m.texts {
		objects = append(objects, m.texts[id].objects...)
	}
	return objects
}

func (m *textManager) New(id string, x, y int) *text {
	m.texts[id] = &text{
		tm: m,
		x:  x,
		y:  y,
	}
	return m.Get(id)
}

func (m *textManager) Get(id string) *text {
	return m.texts[id]
}

func (m *textManager) Delete(id string) {
	delete(m.texts, id)
}

type text struct {
	tm      *textManager
	x       int
	y       int
	s       string
	objects []engine.Object
}

func (t *text) SetContent(s string) {
	chars := strings.Split(t.tm.stringMap(s), "")
	t.objects = make([]engine.Object, len(chars))
	for i := range chars {
		t.objects[i] = engine.Object{
			Key: "char_" + chars[i],
			X:   t.x + i*6,
			Y:   t.y,
		}
	}
}

const allowedChars = "01234567890abcdefghijklmnopqrstuvwxyz.,:;-_()<>!? "

func createFilter(allowedChars string) func(rune) rune {
	m := map[rune]struct{}{}
	for _, r := range []rune(allowedChars) {
		m[r] = struct{}{}
	}
	return func(r rune) rune {
		if _, ok := m[r]; ok {
			return r
		}
		return -1
	}
}
