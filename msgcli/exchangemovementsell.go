// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementSell struct{}

func (m ExchangeMovementSell) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementSell
}

func (m ExchangeMovementSell) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeMovementSell) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
