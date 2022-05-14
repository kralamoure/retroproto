package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterAddError struct {
	Reason rune
}

func NewAccountCharacterAddError(extra string) (AccountCharacterAddError, error) {
	var m AccountCharacterAddError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterAddError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterAddError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterAddError
}

func (m AccountCharacterAddError) MessageName() string {
	return "AccountCharacterAddError"
}

func (m AccountCharacterAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *AccountCharacterAddError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	return nil
}
