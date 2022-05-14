// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationError struct{}

func (m AccountCharacterMigrationError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationError
}

func (m AccountCharacterMigrationError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
