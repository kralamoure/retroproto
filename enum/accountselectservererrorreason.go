package enum

var AccountSelectServerErrorReason = newAccountSelectServerErrorReason()

func newAccountSelectServerErrorReason() *accountSelectServerErrorReason {
	return &accountSelectServerErrorReason{
		Default:                            '?',
		ServerFull:                         'F',
		CantChooseCharacterServerDown:      'd',
		CantChooseCharacterServerFull:      'f',
		CantChooseCharacterShopOtherServer: 's',
		CantSelectThisServer:               'r',
	}
}

type accountSelectServerErrorReason struct {
	Default                            rune
	CantChooseCharacterServerDown      rune
	CantChooseCharacterServerFull      rune
	ServerFull                         rune
	CantChooseCharacterShopOtherServer rune
	CantSelectThisServer               rune
}
