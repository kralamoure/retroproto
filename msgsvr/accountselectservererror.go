package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountSelectServerError struct {
	Reason rune
	Extra  string
}

func (m AccountSelectServerError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountSelectServerError
}

func (m AccountSelectServerError) Serialized() (string, error) {
	return string(m.Reason) + m.Extra, nil
}

func (m *AccountSelectServerError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	if len(extra) < 2 {
		return nil
	}

	m.Extra = extra[1:]

	return nil
}
