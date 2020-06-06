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
	return "", nil
}

func (m *ChatSend) Deserialize(extra string) error {
	return nil
}
