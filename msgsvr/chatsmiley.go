// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ChatSmiley struct{}

func NewChatSmiley(extra string) (ChatSmiley, error) {
	var m ChatSmiley

	if err := m.Deserialize(extra); err != nil {
		return ChatSmiley{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatSmiley) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatSmiley
}

func (m ChatSmiley) MessageName() string {
	return "ChatSmiley"
}

func (m ChatSmiley) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ChatSmiley) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
