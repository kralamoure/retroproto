package enum

var GameCreateType = newGameCreateType()

func newGameCreateType() *gameCreateType {
	return &gameCreateType{
		Solo:  1,
		Fight: 2,
	}
}

type gameCreateType struct {
	Solo  int
	Fight int
}
