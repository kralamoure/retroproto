// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeSellSuccess struct{}

func (m ExchangeSellSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeSellSuccess
}

func (m ExchangeSellSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSellSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
