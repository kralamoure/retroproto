// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountResetCharacter struct{}

func (m AccountResetCharacter) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountResetCharacter
}

func (m AccountResetCharacter) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountResetCharacter) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
