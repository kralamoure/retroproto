// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ChatSend struct{}

func (m ChatSend) ProtocolId() d1proto.MsgCliId {
	return d1proto.ChatSend
}

func (m ChatSend) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ChatSend) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
