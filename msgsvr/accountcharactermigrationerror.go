// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterMigrationError struct{}

func (m AccountCharacterMigrationError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterMigrationError
}

func (m AccountCharacterMigrationError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountCharacterMigrationError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
