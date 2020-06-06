// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementPay struct{}

func (m ExchangeMovementPay) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementPay
}

func (m ExchangeMovementPay) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeMovementPay) Deserialize(extra string) error {
	return nil
}
