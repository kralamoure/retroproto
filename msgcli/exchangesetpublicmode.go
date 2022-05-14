// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeSetPublicMode struct{}

func (m ExchangeSetPublicMode) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeSetPublicMode
}

func (m ExchangeSetPublicMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSetPublicMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
