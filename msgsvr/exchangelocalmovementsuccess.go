// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLocalMovementSuccess struct{}

func (m ExchangeLocalMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalMovementSuccess
}

func (m ExchangeLocalMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
