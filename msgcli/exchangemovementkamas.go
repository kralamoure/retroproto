// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMovementKamas struct{}

func (m ExchangeMovementKamas) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeMovementKamas
}

func (m ExchangeMovementKamas) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMovementKamas) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
