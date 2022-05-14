// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountSetNickname struct{}

func (m AccountSetNickname) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSetNickname
}

func (m AccountSetNickname) MessageName() string {
	return "AccountSetNickname"
}

func (m AccountSetNickname) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountSetNickname) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
