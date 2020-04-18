package yic

import (
	"github.com/GodsBoss/ld46/pkg/ui"
)

func Sprites() map[string]ui.Sprite {
	return map[string]ui.Sprite{
		"bg_title": ui.Sprite{
			X:      400,
			Y:      0,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_playing": ui.Sprite{
			X:      400,
			Y:      300,
			W:      400,
			H:      300,
			Frames: 1,
		},
	}
}
