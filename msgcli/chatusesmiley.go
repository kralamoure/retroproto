// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ChatUseSmiley struct{}

func (m ChatUseSmiley) MessageId() retroproto.MsgCliId {
	return retroproto.ChatUseSmiley
}

func (m ChatUseSmiley) MessageName() string {
	return "ChatUseSmiley"
}

func (m ChatUseSmiley) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ChatUseSmiley) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
