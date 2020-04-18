package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const titleStateID = "title"

type title struct{}

func (t *title) Init() {}

func (t *title) Tick(transitioner engine.Transitioner, ms int) {}
