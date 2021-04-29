// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountSetNickname struct{}

func (m AccountSetNickname) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountSetNickname
}

func (m AccountSetNickname) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountSetNickname) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
