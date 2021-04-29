package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountSendIdentity struct {
	Id string
}

func (m AccountSendIdentity) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountSendIdentity
}

func (m AccountSendIdentity) Serialized() (string, error) {
	return m.Id, nil
}

func (m *AccountSendIdentity) Deserialize(extra string) error {
	m.Id = extra

	return nil
}
