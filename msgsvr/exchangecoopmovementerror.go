// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCoopMovementError struct{}

func (m ExchangeCoopMovementError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCoopMovementError
}

func (m ExchangeCoopMovementError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCoopMovementError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
