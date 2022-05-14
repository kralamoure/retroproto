// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLocalDistantSuccess struct{}

func (m ExchangeLocalDistantSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalDistantSuccess
}

func (m ExchangeLocalDistantSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalDistantSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
