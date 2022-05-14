// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountResetCharacter struct{}

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
