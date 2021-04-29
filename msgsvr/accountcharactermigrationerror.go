// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterMigrationError struct{}

func (m AccountCharacterMigrationError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterMigrationError
}

func (m AccountCharacterMigrationError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountCharacterMigrationError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
