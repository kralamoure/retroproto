package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBuyError struct{}

func (m ExchangeBuyError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBuyError
}

func (m ExchangeBuyError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuyError) Deserialize(extra string) error {
	return nil
}
