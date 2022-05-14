package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeSearchSuccess struct{}

func (m ExchangeSearchSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSearchSuccess
}

func (m ExchangeSearchSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchSuccess) Deserialize(extra string) error {
	return nil
}
