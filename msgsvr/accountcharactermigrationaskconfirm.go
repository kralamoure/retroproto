// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterMigrationAskConfirm struct{}

func (m AccountCharacterMigrationAskConfirm) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterMigrationAskConfirm
}

func (m AccountCharacterMigrationAskConfirm) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountCharacterMigrationAskConfirm) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
