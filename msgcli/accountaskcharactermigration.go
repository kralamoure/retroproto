// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountAskCharacterMigration struct{}

func (m AccountAskCharacterMigration) MessageId() retroproto.MsgCliId {
	return retroproto.AccountAskCharacterMigration
}

func (m AccountAskCharacterMigration) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountAskCharacterMigration) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
