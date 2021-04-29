// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountDeleteCharacterMigration struct{}

func (m AccountDeleteCharacterMigration) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountDeleteCharacterMigration
}

func (m AccountDeleteCharacterMigration) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountDeleteCharacterMigration) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
