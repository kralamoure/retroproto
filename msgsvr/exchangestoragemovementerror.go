// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeStorageMovementError struct{}

func (m ExchangeStorageMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeStorageMovementError
}

func (m ExchangeStorageMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStorageMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
