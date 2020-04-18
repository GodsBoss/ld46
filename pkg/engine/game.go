package engine

type Game struct {
	States map[string]State

	currentStateID string
}

func (game *Game) Tick(ms int) {
	game.currentState().Tick(game, ms)
}

func (game *Game) Transition(nextStateKey string) {
	game.currentStateID = nextStateKey
	game.States[nextStateKey].Init()
}

func (game *Game) CurrentStateID() string {
	return game.currentStateID
}

func (game *Game) currentState() State {
	return game.States[game.currentStateID]
}

type Transitioner interface {
	Transition(nextStateKey string)
}

type State interface {
	// Init will be called every time this state is transitioned to.
	Init()

	// Tick invokes a game tick, given a duration in milliseconds.
	Tick(transitioner Transitioner, ms int)
}
