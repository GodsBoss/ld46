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
	scale   int
	objects []engine.Object
}

func (t *text) SetContent(s string) {
	lines := strings.Split(s, "\n")
	t.objects = make([]engine.Object, 0)
	scale := t.scale
	if scale < 1 {
		scale = 1
	}
	for j := range lines {
		chars := strings.Split(t.tm.stringMap(lines[j]), "")
		for i := range chars {
			t.objects = append(
				t.objects,
				engine.Object{
					Key:   "char_" + chars[i],
					X:     float64(t.x + i*6*scale),
					Y:     float64(t.y + j*6),
					Scale: scale,
				},
			)
		}
	}
}

func (t *text) SetScale(scale int) *text {
	t.scale = scale
	return t
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
