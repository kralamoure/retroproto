package enum

var GameActionChallengeJoinErrorReason = newGameActionChallengeJoinErrorReason()

func newGameActionChallengeJoinErrorReason() *gameActionChallengeJoinErrorReason {
	return &gameActionChallengeJoinErrorReason{
		ChallengeFull:           'c',
		TeamFull:                't',
		TeamDifferentAlignment:  'a',
		Guild:                   'g',
		TooLate:                 'l',
		Mutant:                  'm',
		Map:                     'p',
		NotReady:                'r',
		YouAreBusy:              'o',
		OpponentBusy:            'z',
		MultipleAccounts:        'h',
		NoRights:                'i',
		SubscriptionExpired:     's',
		SubscriptionOut:         'n',
		OpponentSubscriptionOut: 'b',
		TeamClosed:              'f',
		Zombie:                  'd',
	}
}

type gameActionChallengeJoinErrorReason struct {
	ChallengeFull           rune
	TeamFull                rune
	TeamDifferentAlignment  rune
	Guild                   rune
	TooLate                 rune
	Mutant                  rune
	Map                     rune
	NotReady                rune
	YouAreBusy              rune
	OpponentBusy            rune
	MultipleAccounts        rune
	NoRights                rune
	SubscriptionExpired     rune
	SubscriptionOut         rune
	OpponentSubscriptionOut rune
	TeamClosed              rune
	Zombie                  rune
}
