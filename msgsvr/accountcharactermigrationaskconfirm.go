// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationAskConfirm struct{}

func NewAccountCharacterMigrationAskConfirm(extra string) (AccountCharacterMigrationAskConfirm, error) {
	var m AccountCharacterMigrationAskConfirm

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterMigrationAskConfirm{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterMigrationAskConfirm) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationAskConfirm
}

func (m AccountCharacterMigrationAskConfirm) MessageName() string {
	return "AccountCharacterMigrationAskConfirm"
}

func (m AccountCharacterMigrationAskConfirm) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationAskConfirm) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
