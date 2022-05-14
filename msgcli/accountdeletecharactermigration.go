// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountDeleteCharacterMigration struct{}

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
