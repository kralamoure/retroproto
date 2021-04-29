// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeSellError struct{}

func (m ExchangeSellError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeSellError
}

func (m ExchangeSellError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeSellError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
