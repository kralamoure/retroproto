// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterMigrationSuccess struct{}

func NewAccountCharacterMigrationSuccess(extra string) (AccountCharacterMigrationSuccess, error) {
	var m AccountCharacterMigrationSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterMigrationSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterMigrationSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterMigrationSuccess
}

func (m AccountCharacterMigrationSuccess) MessageName() string {
	return "AccountCharacterMigrationSuccess"
}

func (m AccountCharacterMigrationSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterMigrationSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
