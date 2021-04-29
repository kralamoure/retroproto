// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCraftPublicMode struct{}

func (m ExchangeCraftPublicMode) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCraftPublicMode
}

func (m ExchangeCraftPublicMode) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCraftPublicMode) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
