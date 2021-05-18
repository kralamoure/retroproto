package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountLoginError struct {
	Reason rune
	Extra  string
}

func (m AccountLoginError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountLoginError
}

func (m AccountLoginError) Serialized() (string, error) {
	return string(m.Reason) + m.Extra, nil
}

func (m *AccountLoginError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	if len(extra) < 2 {
		return nil
	}

	m.Extra = extra[1:]

	return nil
}
