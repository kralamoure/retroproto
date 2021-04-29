package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type AccountSelectServerPlainSuccess struct {
	Host   string
	Port   string
	Ticket string
}

func (m AccountSelectServerPlainSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountSelectServerPlainSuccess
}

func (m AccountSelectServerPlainSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s:%s;%s", m.Host, m.Port, m.Ticket), nil
}

func (m *AccountSelectServerPlainSuccess) Deserialize(extra string) error {
	if extra == "" {
		return d1encoding.ErrInvalidMsg
	}

	data := strings.Split(extra, ";")
	if len(data) != 2 {
		return d1encoding.ErrInvalidMsg
	}

	addr := strings.Split(data[0], ":")

	m.Host = addr[0]
	if len(addr) >= 2 {
		m.Port = addr[1]
	}

	m.Ticket = data[1]

	return nil
}
