// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCraftPublicMode struct{}

func (m ExchangeCraftPublicMode) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftPublicMode
}

func (m ExchangeCraftPublicMode) MessageName() string {
	return "ExchangeCraftPublicMode"
}

func (m ExchangeCraftPublicMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftPublicMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
