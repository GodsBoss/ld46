package yic

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	"github.com/GodsBoss/ld46/pkg/engine"
)

const hiscoreStateID = "hiscore"

type hiscore struct {
	textManager *textManager
	lists       *hiscoreLists
}

func (h *hiscore) Init() {
	h.textManager = newTextManager()
	h.textManager.New("press_t_for_title", 10, 284).SetContent("Press 't' to return to title")
	h.textManager.New("press_t_for_title", 30, 10).SetScale(2).SetContent("Hall Of Fame")

	levelKeys := make([]string, 0, len(h.lists.hiscoresByLevel))
	for key := range h.lists.hiscoresByLevel {
		levelKeys = append(levelKeys, key)
	}
	sort.Strings(levelKeys)
	line := 0
	for i := range levelKeys {
		h.textManager.New("level_"+levelKeys[i], 15, 25+line*6).SetContent("Level " + levelKeys[i])
		line++
		for j := range h.lists.hiscoresByLevel[levelKeys[i]] {
			tmIDPrefix := "level_" + levelKeys[i] + "_" + strconv.Itoa(j) + "_"
			h.textManager.New(tmIDPrefix+"name", 15, 25+line*6).SetContent(h.lists.hiscoresByLevel[levelKeys[i]][j].Name)
			h.textManager.New(tmIDPrefix+"points", 60, 25+line*6).SetContent(
				leftPad(strconv.Itoa(h.lists.hiscoresByLevel[levelKeys[i]][j].Points), 10),
			)
			line++
		}
		line++
	}
}

func (h *hiscore) Tick(ms int) *engine.Transition {
	return nil
}

func (h *hiscore) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	if event.Key == "t" {
		return engine.NewTransition(titleStateID)
	}
	if event.Key == "c" {
		h.lists.Clear()
		return engine.NewTransition(hiscoreStateID)
	}
	if event.Key == "r" {
		h.lists.Reset()
		return engine.NewTransition(hiscoreStateID)
	}
	return nil
}

func (h *hiscore) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_hiscore",
				X:   0,
				Y:   0,
			},
		},
		"ui": h.textManager.Objects(),
	}
}

const hiscoreListsKey = "ld46_hiscores"

type hiscoreLists struct {
	storage Storage

	hiscoresByLevel map[string]hiscoreListEntries
}

func (lists *hiscoreLists) Load() {
	lists.hiscoresByLevel = defaultHiscoreLists()
	raw, ok := lists.storage.Get(hiscoreListsKey)
	if !ok {
		return
	}
	entries := map[string]hiscoreListEntries{}
	err := json.Unmarshal([]byte(raw), &entries)
	if err != nil {
		return
	}
	lists.hiscoresByLevel = entries
}

func (lists *hiscoreLists) save() {
	data, err := json.Marshal(lists.hiscoresByLevel)
	if err != nil {
		return
	}
	lists.storage.Set(hiscoreListsKey, string(data))
}

func (lists *hiscoreLists) Add(levelKey string, name string, points int) {
	if _, ok := lists.hiscoresByLevel[levelKey]; !ok {
		lists.hiscoresByLevel[levelKey] = hiscoreListEntries{}
	}
	lists.hiscoresByLevel[levelKey] = append(
		lists.hiscoresByLevel[levelKey],
		hiscoreListEntry{
			Name:   name,
			Points: points,
		},
	)
	sort.Sort(lists.hiscoresByLevel[levelKey])
	if len(lists.hiscoresByLevel[levelKey]) > maxHiscoreEntries {
		lists.hiscoresByLevel[levelKey] = lists.hiscoresByLevel[levelKey][:maxHiscoreEntries]
	}
	lists.save()
}

func (lists *hiscoreLists) Clear() {
	lists.hiscoresByLevel = map[string]hiscoreListEntries{}
	lists.save()
}

func (lists *hiscoreLists) Reset() {
	lists.hiscoresByLevel = defaultHiscoreLists()
	lists.save()
}

type hiscoreListEntries []hiscoreListEntry

func (entries hiscoreListEntries) Len() int {
	return len(entries)
}

func (entries hiscoreListEntries) Less(i, j int) bool {
	return entries[i].Points > entries[j].Points
}

func (entries hiscoreListEntries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

type hiscoreListEntry struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func defaultHiscoreLists() map[string]hiscoreListEntries {
	return map[string]hiscoreListEntries{
		"1": hiscoreListEntries{
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 100000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 10000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 1000,
			},
		},
		"2": hiscoreListEntries{
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 100000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 10000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 1000,
			},
		},
		"3": hiscoreListEntries{
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 100000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 10000,
			},
			hiscoreListEntry{
				Name:   "GodsBoss",
				Points: 1000,
			},
		},
	}
}

const maxHiscoreEntries = 4

func leftPad(input string, length int) string {
	if len(input) < length {
		return strings.Repeat(" ", length-len(input)) + input
	}
	return input
}
