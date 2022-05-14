// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationError struct{}

func NewAccountCharacterMigrationError(extra string) (AccountCharacterMigrationError, error) {
	var m AccountCharacterMigrationError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterMigrationError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterMigrationError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationError
}

func (m AccountCharacterMigrationError) MessageName() string {
	return "AccountCharacterMigrationError"
}

func (m AccountCharacterMigrationError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
