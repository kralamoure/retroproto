package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSendTicket struct {
	Ticket string
}

func NewAccountSendTicket(extra string) (AccountSendTicket, error) {
	var m AccountSendTicket

	if err := m.Deserialize(extra); err != nil {
		return AccountSendTicket{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSendTicket) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSendTicket
}

func (m AccountSendTicket) MessageName() string {
	return "AccountSendTicket"
}

func (m AccountSendTicket) Serialized() (string, error) {
	return m.Ticket, nil
}

func (m *AccountSendTicket) Deserialize(extra string) error {
	m.Ticket = extra

	return nil
}
