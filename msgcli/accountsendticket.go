package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountSendTicket struct {
	Ticket string
}

func (m AccountSendTicket) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountSendTicket
}

func (m AccountSendTicket) Serialized() (string, error) {
	return m.Ticket, nil
}

func (m *AccountSendTicket) Deserialize(extra string) error {
	m.Ticket = extra

	return nil
}
