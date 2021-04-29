// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMountPark struct{}

func (m ExchangeMountPark) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeMountPark
}

func (m ExchangeMountPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMountPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
