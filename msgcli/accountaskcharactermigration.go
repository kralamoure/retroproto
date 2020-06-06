// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountAskCharacterMigration struct{}

func (m AccountAskCharacterMigration) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountAskCharacterMigration
}

func (m AccountAskCharacterMigration) Serialized() (string, error) {
	return "", nil
}

func (m *AccountAskCharacterMigration) Deserialize(extra string) error {
	return nil
}
