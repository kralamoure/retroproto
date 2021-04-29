// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeAccept struct{}

func (m ExchangeAccept) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeAccept
}

func (m ExchangeAccept) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeAccept) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
