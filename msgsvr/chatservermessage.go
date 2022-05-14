package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ChatServerMessage struct {
	Message string
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
