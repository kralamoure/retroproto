// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCoopMovementSuccess struct{}

func (m ExchangeCoopMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCoopMovementSuccess
}

func (m ExchangeCoopMovementSuccess) MessageName() string {
	return "ExchangeCoopMovementSuccess"
}

func (m ExchangeCoopMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCoopMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
