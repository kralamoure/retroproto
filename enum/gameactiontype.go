package enum

var GameActionType = newGameActionType()

func newGameActionType() *gameActionType {
	return &gameActionType{
		Default:     0,
		Movement:    1,
		LoadGameMap: 2,
		Challenge:   900,
	}
}

type gameActionType struct {
	Default     int
	Movement    int
	LoadGameMap int
	Challenge   int
}
