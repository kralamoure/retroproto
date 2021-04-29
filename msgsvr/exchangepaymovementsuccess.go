// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePayMovementSuccess struct{}

func (m ExchangePayMovementSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangePayMovementSuccess
}

func (m ExchangePayMovementSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePayMovementSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
