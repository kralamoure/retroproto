// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLocalDistantError struct{}

func (m ExchangeLocalDistantError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalDistantError
}

func (m ExchangeLocalDistantError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalDistantError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
