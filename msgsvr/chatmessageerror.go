package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ChatMessageError struct {
	Reason rune
}

func (m ChatMessageError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ChatMessageError
}

func (m ChatMessageError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ChatMessageError) Deserialize(extra string) error {
	if len(extra) != 1 {
		return d1encoding.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
	}

	return nil
}
