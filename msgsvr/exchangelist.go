// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeList struct{}

func (m ExchangeList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeList
}

func (m ExchangeList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
