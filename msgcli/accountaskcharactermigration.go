// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountAskCharacterMigration struct{}

func NewAccountAskCharacterMigration(extra string) (AccountAskCharacterMigration, error) {
	var m AccountAskCharacterMigration

	if err := m.Deserialize(extra); err != nil {
		return AccountAskCharacterMigration{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountAskCharacterMigration) MessageId() retroproto.MsgCliId {
	return retroproto.AccountAskCharacterMigration
}

func (m AccountAskCharacterMigration) MessageName() string {
	return "AccountAskCharacterMigration"
}

func (m AccountAskCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountAskCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
