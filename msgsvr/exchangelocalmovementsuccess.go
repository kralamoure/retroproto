// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLocalMovementSuccess struct{}

func (m ExchangeLocalMovementSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLocalMovementSuccess
}

func (m ExchangeLocalMovementSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeLocalMovementSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
