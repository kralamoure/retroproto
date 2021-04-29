// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLocalMovementError struct{}

func (m ExchangeLocalMovementError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLocalMovementError
}

func (m ExchangeLocalMovementError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeLocalMovementError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
