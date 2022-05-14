// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLocalMovementError struct{}

func (m ExchangeLocalMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalMovementError
}

func (m ExchangeLocalMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
