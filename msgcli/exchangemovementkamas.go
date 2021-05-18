// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMovementKamas struct{}

func (m ExchangeMovementKamas) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementKamas
}

func (m ExchangeMovementKamas) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementKamas) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
