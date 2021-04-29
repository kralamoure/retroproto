package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatServerMessage struct {
	Message string
}

func (m ChatServerMessage) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatServerMessage
}

func (m ChatServerMessage) Serialized() (string, error) {
	return m.Message, nil
}

func (m *ChatServerMessage) Deserialize(extra string) error {
	m.Message = extra
	return nil
}
