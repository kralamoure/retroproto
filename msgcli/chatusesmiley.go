// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ChatUseSmiley struct{}

func (m ChatUseSmiley) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ChatUseSmiley
}

func (m ChatUseSmiley) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ChatUseSmiley) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
