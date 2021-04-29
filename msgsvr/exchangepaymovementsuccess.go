// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangePayMovementSuccess struct{}

func (m ExchangePayMovementSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangePayMovementSuccess
}

func (m ExchangePayMovementSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangePayMovementSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
