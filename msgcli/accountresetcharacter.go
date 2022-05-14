// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountResetCharacter struct{}

func NewAccountResetCharacter(extra string) (AccountResetCharacter, error) {
	var m AccountResetCharacter

	if err := m.Deserialize(extra); err != nil {
		return AccountResetCharacter{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountResetCharacter) MessageId() retroproto.MsgCliId {
	return retroproto.AccountResetCharacter
}

func (m AccountResetCharacter) MessageName() string {
	return "AccountResetCharacter"
}

func (m AccountResetCharacter) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountResetCharacter) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
