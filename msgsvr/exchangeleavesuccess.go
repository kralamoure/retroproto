package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLeaveSuccess struct {
	TypePlayerExchange bool
}

func (m ExchangeLeaveSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLeaveSuccess
}

func (m ExchangeLeaveSuccess) Serialized() (string, error) {
	if m.TypePlayerExchange {
		return "a", nil
	}
	return "", nil
}

func (m *ExchangeLeaveSuccess) Deserialize(extra string) error {
	if extra == "a" {
		m.TypePlayerExchange = true
	}
	return nil
}
