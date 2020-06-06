package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountSendTicket struct {
	Id string
}

func (m AccountSendTicket) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountSendTicket
}

func (m AccountSendTicket) Serialized() (string, error) {
	return m.Id, nil
}

func (m *AccountSendTicket) Deserialize(extra string) error {
	m.Id = extra

	return nil
}
