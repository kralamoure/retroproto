// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountValidCharacterMigration struct{}

func (m AccountValidCharacterMigration) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountValidCharacterMigration
}

func (m AccountValidCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountValidCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
