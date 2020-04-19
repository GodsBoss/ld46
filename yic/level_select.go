package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"

	"sort"
)

const levelSelectStateID = "level_select"

type levelSelect struct {
	levels *levels

	textManager *textManager
}

func (ls *levelSelect) Init() {
	ls.textManager = newTextManager()
	ls.textManager.New("label_choose_level", 20, 40).SetContent("Choose level (press corresponding key):")
	keys := make([]string, 0, len(ls.levels.byKey))
	for key := range ls.levels.byKey {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i := range keys {
		ls.textManager.New("level-"+keys[i], 20, 52+6*i).SetContent(keys[i])
	}
}

func (ls *levelSelect) Tick(ms int) *engine.Transition {
	return nil
}

func (ls *levelSelect) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	switch event.Key {
	case "b":
		return engine.NewTransition(titleStateID)
	}
	if _, ok := ls.levels.byKey[event.Key]; ok {
		ls.levels.chosen = event.Key
		return engine.NewTransition(playingStateID)
	}
	return nil
}

func (ls *levelSelect) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_level_select",
				X:   0,
				Y:   0,
			},
		},
		"ui": ls.textManager.Objects(),
	}
}
