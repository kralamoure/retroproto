// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCoopMovementSuccess struct{}

func (m ExchangeCoopMovementSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCoopMovementSuccess
}

func (m ExchangeCoopMovementSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCoopMovementSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
