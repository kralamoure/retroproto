package enum

var AccountCharacterAddErrorReason = newAccountCharacterAddErrorReason()

func newAccountCharacterAddErrorReason() *accountCharacterAddErrorReason {
	return &accountCharacterAddErrorReason{
		Default:                '?',
		SubscriptionOut:        's',
		CreateCharacterFull:    'f',
		NameAlreadyExists:      'a',
		CreateCharacterBadName: 'n',
	}
}

type accountCharacterAddErrorReason struct {
	Default                rune
	SubscriptionOut        rune
	CreateCharacterFull    rune
	NameAlreadyExists      rune
	CreateCharacterBadName rune
}
