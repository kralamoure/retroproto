// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeStorageMovementSuccess struct{}

func (m ExchangeStorageMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeStorageMovementSuccess
}

func (m ExchangeStorageMovementSuccess) MessageName() string {
	return "ExchangeStorageMovementSuccess"
}

func (m ExchangeStorageMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStorageMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
