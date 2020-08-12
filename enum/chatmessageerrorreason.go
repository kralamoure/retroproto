package enum

var ChatMessageErrorReason = newChatMessageErrorReason()

func newChatMessageErrorReason() *chatMessageErrorReason {
	return &chatMessageErrorReason{
		Syntax:                         'S',
		NotConnected:                   'f',
		NotConnectedButTrySendExternal: 'e',
		NotConnectedExternalNack:       'n',
	}
}

type chatMessageErrorReason struct {
	Syntax                         rune
	NotConnected                   rune
	NotConnectedButTrySendExternal rune
	NotConnectedExternalNack       rune
}
