package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ChatServerMessage struct {
	Message string
}

func NewChatServerMessage(extra string) (ChatServerMessage, error) {
	var m ChatServerMessage

	if err := m.Deserialize(extra); err != nil {
		return ChatServerMessage{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatServerMessage) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatServerMessage
}

func (m ChatServerMessage) MessageName() string {
	return "ChatServerMessage"
}

func (m ChatServerMessage) Serialized() (string, error) {
	return m.Message, nil
}

func (m *ChatServerMessage) Deserialize(extra string) error {
	m.Message = extra
	return nil
}
