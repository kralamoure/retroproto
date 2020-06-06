package enum

var GameActionType = newGameActionType()

func newGameActionType() *gameActionType {
	return &gameActionType{
		Default:     0,
		Movement:    1,
		LoadGameMap: 2,
	}
}

type gameActionType struct {
	Default     int
	Movement    int
	LoadGameMap int
}
