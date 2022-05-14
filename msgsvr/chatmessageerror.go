package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ChatMessageError struct {
	Reason rune
}

func NewChatMessageError(extra string) (ChatMessageError, error) {
	var m ChatMessageError

	if err := m.Deserialize(extra); err != nil {
		return ChatMessageError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatMessageError) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatMessageError
}

func (m ChatMessageError) MessageName() string {
	return "ChatMessageError"
}

func (m ChatMessageError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ChatMessageError) Deserialize(extra string) error {
	if len(extra) != 1 {
		return retroproto.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
	}

	return nil
}
