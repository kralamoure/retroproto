// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterMigrationAskConfirm struct{}

func (m AccountCharacterMigrationAskConfirm) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterMigrationAskConfirm
}

func (m AccountCharacterMigrationAskConfirm) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountCharacterMigrationAskConfirm) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
