// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCraftPublicMode struct{}

func (m ExchangeCraftPublicMode) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCraftPublicMode
}

func (m ExchangeCraftPublicMode) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCraftPublicMode) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
