// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountValidCharacterMigration struct{}

func (m AccountValidCharacterMigration) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountValidCharacterMigration
}

func (m AccountValidCharacterMigration) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountValidCharacterMigration) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
