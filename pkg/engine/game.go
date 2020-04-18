package engine

type Game struct {
	States map[string]State

	currentStateID string
}

func (game *Game) Tick(ms int) {
	game.currentState().Tick(ms).invoke(game)
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

type Transition struct {
	NextStateKey string
}

func NewTransition(nextStateKey string) *Transition {
	return &Transition{
		NextStateKey: nextStateKey,
	}
}

func (transition *Transition) invoke(game *Game) {
	if transition != nil {
		game.Transition(transition.NextStateKey)
	}
}

type State interface {
	// Init will be called every time this state is transitioned to.
	Init()

	// Tick invokes a game tick, given a duration in milliseconds.
	Tick(ms int) *Transition
}
