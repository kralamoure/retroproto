// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountAskCharacterMigration struct{}

func (m AccountAskCharacterMigration) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountAskCharacterMigration
}

func (m AccountAskCharacterMigration) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountAskCharacterMigration) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
