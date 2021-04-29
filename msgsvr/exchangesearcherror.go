package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeSearchError struct{}

func (m ExchangeSearchError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeSearchError
}

func (m ExchangeSearchError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchError) Deserialize(extra string) error {
	return nil
}
