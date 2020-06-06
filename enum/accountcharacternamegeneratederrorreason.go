package enum

var AccountCharacterNameGeneratedErrorReason = newAccountCharacterNameGeneratedErrorReason()

func newAccountCharacterNameGeneratedErrorReason() *accountCharacterNameGeneratedErrorReason {
	return &accountCharacterNameGeneratedErrorReason{
		Default:           1,
		CantGenerateNames: 2,
	}
}

type accountCharacterNameGeneratedErrorReason struct {
	Default           int
	CantGenerateNames int
}
