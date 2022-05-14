package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterSelectedError struct{}

func NewAccountCharacterSelectedError(extra string) (AccountCharacterSelectedError, error) {
	var m AccountCharacterSelectedError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterSelectedError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterSelectedError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterSelectedError
}

func (m AccountCharacterSelectedError) MessageName() string {
	return "AccountCharacterSelectedError"
}

func (m AccountCharacterSelectedError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterSelectedError) Deserialize(extra string) error {
	return nil
}
