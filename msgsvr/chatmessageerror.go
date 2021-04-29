package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ChatMessageError struct {
	Reason rune
}

func (m ChatMessageError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatMessageError
}

func (m ChatMessageError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ChatMessageError) Deserialize(extra string) error {
	if len(extra) != 1 {
		return d1proto.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
	}

	return nil
}
