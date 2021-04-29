// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ChatSmiley struct{}

func (m ChatSmiley) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ChatSmiley
}

func (m ChatSmiley) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ChatSmiley) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
