package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLeaveError struct{}

func (m ExchangeLeaveError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLeaveError
}

func (m ExchangeLeaveError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLeaveError) Deserialize(extra string) error {
	return nil
}
