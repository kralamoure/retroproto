package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterDeleteError struct{}

func NewAccountCharacterDeleteError(extra string) (AccountCharacterDeleteError, error) {
	var m AccountCharacterDeleteError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterDeleteError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterDeleteError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterDeleteError
}

func (m AccountCharacterDeleteError) MessageName() string {
	return "AccountCharacterDeleteError"
}

func (m AccountCharacterDeleteError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterDeleteError) Deserialize(extra string) error {
	return nil
}
