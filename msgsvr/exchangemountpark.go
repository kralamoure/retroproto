// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMountPark struct{}

func (m ExchangeMountPark) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountPark
}

func (m ExchangeMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
