// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLocalMovementError struct{}

func (m ExchangeLocalMovementError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLocalMovementError
}

func (m ExchangeLocalMovementError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeLocalMovementError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
