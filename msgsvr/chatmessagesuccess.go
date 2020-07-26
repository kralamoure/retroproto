// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatMessageSuccess struct{}

func (m ChatMessageSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageSuccess
}

func (m ChatMessageSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ChatMessageSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
