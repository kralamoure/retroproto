// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ChatUseSmiley struct{}

func (m ChatUseSmiley) ProtocolId() d1proto.MsgCliId {
	return d1proto.ChatUseSmiley
}

func (m ChatUseSmiley) Serialized() (string, error) {
	return "", nil
}

func (m *ChatUseSmiley) Deserialize(extra string) error {
	return nil
}
