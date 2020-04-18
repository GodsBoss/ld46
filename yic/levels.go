package yic

type levels struct {
	byKey  map[string]level
	chosen string
}

type level struct{}

func createLevels() *levels {
	return &levels{
		byKey: map[string]level{
			"1": level{},
		},
	}
}
