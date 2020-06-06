// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterMigrationSuccess struct{}

func (m AccountCharacterMigrationSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterMigrationSuccess
}

func (m AccountCharacterMigrationSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterMigrationSuccess) Deserialize(extra string) error {
	return nil
}
