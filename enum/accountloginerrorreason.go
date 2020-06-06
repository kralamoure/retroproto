package enum

var AccountLoginErrorReason = newAccountLoginErrorReason()

func newAccountLoginErrorReason() *accountLoginErrorReason {
	return &accountLoginErrorReason{
		Default:                   '?',
		ConnectNotFinished:        'n',
		AlreadyLogged:             'a',
		AlreadyLoggedGameServer:   'c',
		BadVersion:                'v',
		NotPlayer:                 'p',
		Banned:                    'b',
		UDisconnectAccount:        'd',
		Kicked:                    'k',
		ServerFull:                'w',
		OldAccount:                'o',
		OldAccountUseNew:          'e',
		MaintainAccount:           'm',
		ChooseNickname:            'r',
		ChooseNicknameAlreadyUsed: 's',
		AccessDenied:              'f',
	}
}

type accountLoginErrorReason struct {
	Default                   rune
	ConnectNotFinished        rune
	AlreadyLogged             rune
	AlreadyLoggedGameServer   rune
	BadVersion                rune
	NotPlayer                 rune
	Banned                    rune
	UDisconnectAccount        rune
	Kicked                    rune
	ServerFull                rune
	OldAccount                rune
	OldAccountUseNew          rune
	MaintainAccount           rune
	ChooseNickname            rune
	ChooseNicknameAlreadyUsed rune
	AccessDenied              rune
}
