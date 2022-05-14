// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationSuccess struct{}

func (m AccountCharacterMigrationSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationSuccess
}

func (m AccountCharacterMigrationSuccess) MessageName() string {
	return "AccountCharacterMigrationSuccess"
}

func (m AccountCharacterMigrationSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
