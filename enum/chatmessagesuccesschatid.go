package enum

var ChatMessageSuccessChatId = newChatMessageSuccessChatId()

func newChatMessageSuccessChatId() *chatMessageSuccessChatId {
	return &chatMessageSuccessChatId{
		Public:       "",
		SecretEmotes: "F",
		Private:      "T",
		Team:         "#",
		Guild:        "%",
		Group:        "$",
		Alignment:    "!",
		Recruitment:  "?",
		Trading:      ":",
		Newbies:      "^",
		Admin:        "@",
	}
}

type chatMessageSuccessChatId struct {
	Public       string
	SecretEmotes string
	Private      string
	Team         string
	Guild        string
	Group        string
	Alignment    string
	Recruitment  string
	Trading      string
	Newbies      string
	Admin        string
}
