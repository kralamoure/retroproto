package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSelectServerError struct {
	Reason rune
	Extra  string
}

func NewAccountSelectServerError(extra string) (AccountSelectServerError, error) {
	var m AccountSelectServerError

	if err := m.Deserialize(extra); err != nil {
		return AccountSelectServerError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSelectServerError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountSelectServerError
}

func (m AccountSelectServerError) MessageName() string {
	return "AccountSelectServerError"
}

func (m AccountSelectServerError) Serialized() (string, error) {
	return string(m.Reason) + m.Extra, nil
}

func (m *AccountSelectServerError) Deserialize(extra string) error {
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
