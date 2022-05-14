// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeAccept struct{}

func (m ExchangeAccept) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeAccept
}

func (m ExchangeAccept) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeAccept) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
