package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCreateError struct{}

func (m ExchangeCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCreateError
}

func (m ExchangeCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCreateError) Deserialize(extra string) error {
	return nil
}
