package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type InfosMessage struct {
	ChatId   int
	Messages []typ.InfosMessageMessage
}

func (m InfosMessage) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosMessage
}

func (m InfosMessage) MessageName() string {
	return "InfosMessage"
}

func (m InfosMessage) Serialized() (string, error) {
	messages := make([]string, len(m.Messages))
	for i, v := range m.Messages {
		message, err := v.Serialized()
		if err != nil {
			return "", err
		}
		messages[i] = message
	}

	return fmt.Sprintf("%d%s", m.ChatId, strings.Join(messages, "|")), nil
}

func (m *InfosMessage) Deserialize(extra string) error {
	if len(extra) < 2 {
		return retroproto.ErrInvalidMsg
	}

	chatID, err := strconv.ParseInt(extra[:1], 10, 32)
	if err != nil {
		return err
	}
	m.ChatId = int(chatID)

	sli := strings.Split(extra[1:], "|")

	m.Messages = make([]typ.InfosMessageMessage, len(sli))
	for i, v := range sli {
		message := &typ.InfosMessageMessage{}
		err := message.Deserialize(v)
		if err != nil {
			return err
		}
		m.Messages[i] = *message
	}

	return nil
}
