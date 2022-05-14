package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountSendIdentity struct {
	Id string
}

func (m AccountSendIdentity) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSendIdentity
}

func (m AccountSendIdentity) Serialized() (string, error) {
	return m.Id, nil
}

func (m *AccountSendIdentity) Deserialize(extra string) error {
	m.Id = extra

	return nil
}
