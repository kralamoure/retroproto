// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeRequestSuccess struct{}

func (m ExchangeRequestSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeRequestSuccess
}

func (m ExchangeRequestSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeRequestSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
