package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLeaveSuccess struct {
	TypePlayerExchange bool
}

func (m ExchangeLeaveSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLeaveSuccess
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
