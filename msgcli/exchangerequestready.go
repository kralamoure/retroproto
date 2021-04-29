// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeRequestReady struct{}

func (m ExchangeRequestReady) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeRequestReady
}

func (m ExchangeRequestReady) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeRequestReady) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
