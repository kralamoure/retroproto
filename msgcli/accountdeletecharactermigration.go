// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountDeleteCharacterMigration struct{}

func NewAccountDeleteCharacterMigration(extra string) (AccountDeleteCharacterMigration, error) {
	var m AccountDeleteCharacterMigration

	if err := m.Deserialize(extra); err != nil {
		return AccountDeleteCharacterMigration{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountDeleteCharacterMigration) MessageId() retroproto.MsgCliId {
	return retroproto.AccountDeleteCharacterMigration
}

func (m AccountDeleteCharacterMigration) MessageName() string {
	return "AccountDeleteCharacterMigration"
}

func (m AccountDeleteCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountDeleteCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
