// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountValidCharacterMigration struct{}

func (m AccountValidCharacterMigration) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountValidCharacterMigration
}

func (m AccountValidCharacterMigration) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountValidCharacterMigration) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
