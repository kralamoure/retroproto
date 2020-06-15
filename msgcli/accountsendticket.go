package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountSendTicket struct {
	Ticket string
}

func (m AccountSendTicket) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountSendTicket
}

func (m AccountSendTicket) Serialized() (string, error) {
	return m.Ticket, nil
}

func (m *AccountSendTicket) Deserialize(extra string) error {
	m.Ticket = extra

	return nil
}
