// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ChatUseSmiley struct{}

func NewChatUseSmiley(extra string) (ChatUseSmiley, error) {
	var m ChatUseSmiley

	if err := m.Deserialize(extra); err != nil {
		return ChatUseSmiley{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
