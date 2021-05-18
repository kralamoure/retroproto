package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCreateError struct{}

func (m ExchangeCreateError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeCreateError
}

func (m ExchangeCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCreateError) Deserialize(extra string) error {
	return nil
}
