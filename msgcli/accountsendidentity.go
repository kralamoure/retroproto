package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountSendIdentity struct {
	Id string
}

func (m AccountSendIdentity) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountSendIdentity
}

func (m AccountSendIdentity) Serialized() (string, error) {
	return m.Id, nil
}

func (m *AccountSendIdentity) Deserialize(extra string) error {
	m.Id = extra

	return nil
}
