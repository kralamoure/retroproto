package msgsvr

import (
	"fmt"

	"github.com/kralamoure/d1proto"
)

type AccountSelectServerSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func (m AccountSelectServerSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountSelectServerSuccess
}

// TODO
func (m AccountSelectServerSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountSelectServerSuccess) Deserialize(extra string) error {
	host, port, ticket, err := d1proto.SplitEncodedHostPortTicket(extra)
	if err != nil {
		return err
	}
	m.Host = host
	m.Port = fmt.Sprint(port)
	m.Ticket = ticket

	return nil
}
