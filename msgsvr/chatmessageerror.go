// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatMessageError struct{}

func (m ChatMessageError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageError
}

func (m ChatMessageError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ChatMessageError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
