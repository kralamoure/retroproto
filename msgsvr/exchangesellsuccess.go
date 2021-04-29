// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeSellSuccess struct{}

func (m ExchangeSellSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeSellSuccess
}

func (m ExchangeSellSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeSellSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
