// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeStorageMovementSuccess struct{}

func (m ExchangeStorageMovementSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeStorageMovementSuccess
}

func (m ExchangeStorageMovementSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeStorageMovementSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
