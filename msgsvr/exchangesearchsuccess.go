package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeSearchSuccess struct{}

func (m ExchangeSearchSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeSearchSuccess
}

func (m ExchangeSearchSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchSuccess) Deserialize(extra string) error {
	return nil
}
