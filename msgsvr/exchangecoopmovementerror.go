// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCoopMovementError struct{}

func (m ExchangeCoopMovementError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCoopMovementError
}

func (m ExchangeCoopMovementError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCoopMovementError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
