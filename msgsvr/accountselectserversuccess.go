package msgsvr

import (
	"fmt"

	"github.com/kralamoure/d1encoding"
)

type AccountSelectServerSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func (m AccountSelectServerSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountSelectServerSuccess
}

// TODO
func (m AccountSelectServerSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountSelectServerSuccess) Deserialize(extra string) error {
	if extra == "" {
		return d1encoding.ErrInvalidMsg
	}

	host, port, ticket, err := d1encoding.SplitEncodedHostPortTicket(extra)
	if err != nil {
		return err
	}
	m.Host = host
	m.Port = fmt.Sprint(port)
	m.Ticket = ticket

	return nil
}
