package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSelectServerSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func NewAccountSelectServerSuccess(extra string) (AccountSelectServerSuccess, error) {
	var m AccountSelectServerSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountSelectServerSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSelectServerSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountSelectServerSuccess
}

func (m AccountSelectServerSuccess) MessageName() string {
	return "AccountSelectServerSuccess"
}

// TODO
func (m AccountSelectServerSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountSelectServerSuccess) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
	}

	host, port, ticket, err := retroproto.SplitEncodedHostPortTicket(extra)
	if err != nil {
		return err
	}
	m.Host = host
	m.Port = fmt.Sprint(port)
	m.Ticket = ticket

	return nil
}
