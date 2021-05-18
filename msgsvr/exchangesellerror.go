// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeSellError struct{}

func (m ExchangeSellError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeSellError
}

func (m ExchangeSellError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSellError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
