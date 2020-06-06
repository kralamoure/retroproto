// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeMovementKamas struct{}

func (m ExchangeMovementKamas) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeMovementKamas
}

func (m ExchangeMovementKamas) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeMovementKamas) Deserialize(extra string) error {
	return nil
}
