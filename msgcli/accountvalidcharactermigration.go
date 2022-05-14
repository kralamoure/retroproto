// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountValidCharacterMigration struct{}

func NewAccountValidCharacterMigration(extra string) (AccountValidCharacterMigration, error) {
	var m AccountValidCharacterMigration

	if err := m.Deserialize(extra); err != nil {
		return AccountValidCharacterMigration{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountValidCharacterMigration) MessageId() retroproto.MsgCliId {
	return retroproto.AccountValidCharacterMigration
}

func (m AccountValidCharacterMigration) MessageName() string {
	return "AccountValidCharacterMigration"
}

func (m AccountValidCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountValidCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
