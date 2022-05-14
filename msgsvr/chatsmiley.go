// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ChatSmiley struct{}

func (m ChatSmiley) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatSmiley
}

func (m ChatSmiley) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ChatSmiley) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
