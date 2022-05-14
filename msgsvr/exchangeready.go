// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeReady struct{}

func (m ExchangeReady) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeReady
}

func (m ExchangeReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
