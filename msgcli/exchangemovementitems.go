// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMovementItems struct{}

func (m ExchangeMovementItems) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeMovementItems
}

func (m ExchangeMovementItems) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMovementItems) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
