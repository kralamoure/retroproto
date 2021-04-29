package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCreateError struct{}

func (m ExchangeCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCreateError
}

func (m ExchangeCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCreateError) Deserialize(extra string) error {
	return nil
}
