package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountSelectServerError struct {
	Reason rune
	Extra  string
}

func (m AccountSelectServerError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountSelectServerError
}

func (m AccountSelectServerError) Serialized() (string, error) {
	return string(m.Reason) + m.Extra, nil
}

func (m *AccountSelectServerError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1encoding.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	if len(extra) < 2 {
		return nil
	}

	m.Extra = extra[1:]

	return nil
}
