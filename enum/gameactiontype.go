package enum

var GameActionType = newGameActionType()

func newGameActionType() *gameActionType {
	return &gameActionType{
		Default:         0,
		Movement:        1,
		LoadGameMap:     2,
		Challenge:       900,
		ChallengeAccept: 901,
		ChallengeRefuse: 902,
		ChallengeJoin:   903,
	}
}

type gameActionType struct {
	Default         int
	Movement        int
	LoadGameMap     int
	Challenge       int
	ChallengeAccept int
	ChallengeRefuse int
	ChallengeJoin   int
}
