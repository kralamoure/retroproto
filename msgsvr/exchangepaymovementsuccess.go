// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangePayMovementSuccess struct{}

func (m ExchangePayMovementSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangePayMovementSuccess
}

func (m ExchangePayMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePayMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
