// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterMigrationSuccess struct{}

func (m AccountCharacterMigrationSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterMigrationSuccess
}

func (m AccountCharacterMigrationSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountCharacterMigrationSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
