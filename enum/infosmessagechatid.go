package enum

var InfosMessageChatId = newInfosMessageChatId()

func newInfosMessageChatId() *infosMessageChatId {
	return &infosMessageChatId{
		Info:  0,
		Error: 1,
		PVP:   2,
	}
}

type infosMessageChatId struct {
	Info  int
	Error int
	PVP   int
}
