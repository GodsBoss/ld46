package yic

type incomeBuilding struct {
	pos vector2D
}

func (b *incomeBuilding) gridXY() vector2D {
	return b.pos
}

func (b *incomeBuilding) typ() string {
	return "building_income"
}

func (b *incomeBuilding) IncomePerSecond() float64 {
	return 15.0
}

// Ensure incomeBuilding fulfills building.
var _ building = &incomeBuilding{}
