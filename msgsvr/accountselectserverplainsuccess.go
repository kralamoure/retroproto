package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountSelectServerPlainSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func (m AccountSelectServerPlainSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountSelectServerPlainSuccess
}

func (m AccountSelectServerPlainSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s:%s;%s", m.Host, m.Port, m.Ticket), nil
}

func (m *AccountSelectServerPlainSuccess) Deserialize(extra string) error {
	data := strings.Split(extra, ";")
	addr := strings.Split(data[0], ":")

	m.Host = addr[0]
	m.Port = addr[1]
	m.Ticket = data[1]

	return nil
}
