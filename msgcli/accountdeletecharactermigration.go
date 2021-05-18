// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountDeleteCharacterMigration struct{}

func (m AccountDeleteCharacterMigration) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountDeleteCharacterMigration
}

func (m AccountDeleteCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountDeleteCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
