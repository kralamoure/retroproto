// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatServerMessage struct{}

func (m ChatServerMessage) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatServerMessage
}

func (m ChatServerMessage) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ChatServerMessage) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
