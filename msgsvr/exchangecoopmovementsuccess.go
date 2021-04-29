// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCoopMovementSuccess struct{}

func (m ExchangeCoopMovementSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCoopMovementSuccess
}

func (m ExchangeCoopMovementSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCoopMovementSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
