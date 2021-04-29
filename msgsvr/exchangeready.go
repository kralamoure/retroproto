// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeReady struct{}

func (m ExchangeReady) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeReady
}

func (m ExchangeReady) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeReady) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
