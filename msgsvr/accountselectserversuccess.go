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

func (m AccountSelectServerSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountSelectServerSuccess
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
