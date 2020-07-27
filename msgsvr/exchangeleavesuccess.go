package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLeaveSuccess struct {
	TypePlayerExchange bool
}

func (m ExchangeLeaveSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLeaveSuccess
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
