// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeStorageMovementSuccess struct{}

func (m ExchangeStorageMovementSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeStorageMovementSuccess
}

func (m ExchangeStorageMovementSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeStorageMovementSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
