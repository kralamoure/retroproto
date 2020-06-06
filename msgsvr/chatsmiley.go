// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatSmiley struct{}

func (m ChatSmiley) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatSmiley
}

func (m ChatSmiley) Serialized() (string, error) {
	return "", nil
}

func (m *ChatSmiley) Deserialize(extra string) error {
	return nil
}
