package yic

import (
	"github.com/GodsBoss/ld46/pkg/ui"
)

func Sprites() map[string]ui.Sprite {
	return map[string]ui.Sprite{
		"field_way": ui.Sprite{
			X:      0,
			Y:      36,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"field_buildspot": ui.Sprite{
			X:      0,
			Y:      0,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"field_obstacle": ui.Sprite{
			X:      0,
			Y:      18,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"head_toddler": ui.Sprite{
			X:      72,
			Y:      0,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"head_child": ui.Sprite{
			X:      72,
			Y:      36,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"head_teen": ui.Sprite{
			X:      72,
			Y:      72,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
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
		"bg_level_select": ui.Sprite{
			X:      400,
			Y:      600,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_game_over": ui.Sprite{
			X:      400,
			Y:      900,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_hiscore": ui.Sprite{
			X:      400,
			Y:      1200,
			W:      400,
			H:      300,
			Frames: 1,
		},
		responsibilityType1: ui.Sprite{
			X:      0,
			Y:      60,
			W:      16,
			H:      16,
			Frames: 4,
		},
		responsibilityType2: ui.Sprite{
			X:      0,
			Y:      76,
			W:      16,
			H:      16,
			Frames: 4,
		},
		responsibilityType3: ui.Sprite{
			X:      0,
			Y:      92,
			W:      16,
			H:      16,
			Frames: 4,
		},
		"grid_cursor": ui.Sprite{
			X:      18,
			Y:      0,
			W:      18,
			H:      18,
			Frames: 1,
		},
	}
}
