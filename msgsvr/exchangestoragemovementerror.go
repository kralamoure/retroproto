// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeStorageMovementError struct{}

func (m ExchangeStorageMovementError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeStorageMovementError
}

func (m ExchangeStorageMovementError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeStorageMovementError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
