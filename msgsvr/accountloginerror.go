package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountLoginError struct {
	Reason rune
	Extra  string
}

func NewAccountLoginError(extra string) (AccountLoginError, error) {
	var m AccountLoginError

	if err := m.Deserialize(extra); err != nil {
		return AccountLoginError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountLoginError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountLoginError
}

func (m AccountLoginError) MessageName() string {
	return "AccountLoginError"
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
