// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeStorageMovementSuccess struct{}

func NewExchangeStorageMovementSuccess(extra string) (ExchangeStorageMovementSuccess, error) {
	var m ExchangeStorageMovementSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeStorageMovementSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
