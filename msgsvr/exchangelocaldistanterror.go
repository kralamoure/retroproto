// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLocalDistantError struct{}

func (m ExchangeLocalDistantError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLocalDistantError
}

func (m ExchangeLocalDistantError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeLocalDistantError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
