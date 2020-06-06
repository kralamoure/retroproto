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
	return "", nil
}

func (m *ChatServerMessage) Deserialize(extra string) error {
	return nil
}
