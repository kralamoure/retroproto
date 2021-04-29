// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountResetCharacter struct{}

func (m AccountResetCharacter) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountResetCharacter
}

func (m AccountResetCharacter) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountResetCharacter) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
