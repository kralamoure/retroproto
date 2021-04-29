// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeStorageMovementError struct{}

func (m ExchangeStorageMovementError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeStorageMovementError
}

func (m ExchangeStorageMovementError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeStorageMovementError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
