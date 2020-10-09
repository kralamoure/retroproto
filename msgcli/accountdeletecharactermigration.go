// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountDeleteCharacterMigration struct{}

func (m AccountDeleteCharacterMigration) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountDeleteCharacterMigration
}

func (m AccountDeleteCharacterMigration) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountDeleteCharacterMigration) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
