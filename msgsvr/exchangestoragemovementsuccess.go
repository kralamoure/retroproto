// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeStorageMovementSuccess struct{}

func (m ExchangeStorageMovementSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeStorageMovementSuccess
}

func (m ExchangeStorageMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStorageMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
