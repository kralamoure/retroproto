package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ChatServerMessage struct {
	Message string
}

func (m ChatServerMessage) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ChatServerMessage
}

func (m ChatServerMessage) Serialized() (string, error) {
	return m.Message, nil
}

func (m *ChatServerMessage) Deserialize(extra string) error {
	m.Message = extra
	return nil
}
