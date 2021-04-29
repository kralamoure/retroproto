// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementItems struct{}

func (m ExchangeMovementItems) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementItems
}

func (m ExchangeMovementItems) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeMovementItems) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
