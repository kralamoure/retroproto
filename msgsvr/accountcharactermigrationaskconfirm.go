// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationAskConfirm struct{}

func (m AccountCharacterMigrationAskConfirm) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationAskConfirm
}

func (m AccountCharacterMigrationAskConfirm) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationAskConfirm) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
