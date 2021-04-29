// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLocalDistantSuccess struct{}

func (m ExchangeLocalDistantSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeLocalDistantSuccess
}

func (m ExchangeLocalDistantSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeLocalDistantSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
