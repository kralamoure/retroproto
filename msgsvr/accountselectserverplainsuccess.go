package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountSelectServerPlainSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func (m AccountSelectServerPlainSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountSelectServerPlainSuccess
}

func (m AccountSelectServerPlainSuccess) MessageName() string {
	return "AccountSelectServerPlainSuccess"
}

func (m AccountSelectServerPlainSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s:%s;%s", m.Host, m.Port, m.Ticket), nil
}

func (m *AccountSelectServerPlainSuccess) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
	}

	data := strings.Split(extra, ";")
	if len(data) != 2 {
		return retroproto.ErrInvalidMsg
	}

	addr := strings.Split(data[0], ":")

	m.Host = addr[0]
	if len(addr) >= 2 {
		m.Port = addr[1]
	}

	m.Ticket = data[1]

	return nil
}
