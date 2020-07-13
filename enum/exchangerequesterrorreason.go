package enum

var ExchangeRequestErrorReason = newExchangeRequestErrorReason()

func newExchangeRequestErrorReason() *exchangeRequestErrorReason {
	return &exchangeRequestErrorReason{
		Default:           'I',
		AlreadyExchange:   'O',
		NotNearCraftTable: 'T',
		MustEquipTool:     'J',
		Overload:          'o',
		Unsubscribed:      'S',
	}
}

type exchangeRequestErrorReason struct {
	Default           rune
	AlreadyExchange   rune
	NotNearCraftTable rune
	MustEquipTool     rune
	Overload          rune
	Unsubscribed      rune
}
