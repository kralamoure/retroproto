// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCoopMovementError struct{}

func (m ExchangeCoopMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCoopMovementError
}

func (m ExchangeCoopMovementError) MessageName() string {
	return "ExchangeCoopMovementError"
}

func (m ExchangeCoopMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCoopMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
