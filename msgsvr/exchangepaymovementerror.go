// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePayMovementError struct{}

func (m ExchangePayMovementError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangePayMovementError
}

func (m ExchangePayMovementError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePayMovementError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
