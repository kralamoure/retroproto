// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLocalMovementSuccess struct{}

func (m ExchangeLocalMovementSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLocalMovementSuccess
}

func (m ExchangeLocalMovementSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeLocalMovementSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
