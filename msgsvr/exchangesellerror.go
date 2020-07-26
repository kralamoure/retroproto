// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeSellError struct{}

func (m ExchangeSellError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeSellError
}

func (m ExchangeSellError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeSellError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
