// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangePayMovementError struct{}

func (m ExchangePayMovementError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangePayMovementError
}

func (m ExchangePayMovementError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangePayMovementError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
