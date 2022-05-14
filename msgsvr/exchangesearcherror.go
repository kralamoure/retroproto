package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeSearchError struct{}

func (m ExchangeSearchError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSearchError
}

func (m ExchangeSearchError) MessageName() string {
	return "ExchangeSearchError"
}

func (m ExchangeSearchError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchError) Deserialize(extra string) error {
	return nil
}
