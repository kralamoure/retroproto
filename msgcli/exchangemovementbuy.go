// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementBuy struct{}

func (m ExchangeMovementBuy) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementBuy
}

func (m ExchangeMovementBuy) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeMovementBuy) Deserialize(extra string) error {
	return nil
}
