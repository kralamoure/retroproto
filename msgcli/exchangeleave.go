// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeLeave struct{}

func (m ExchangeLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeLeave
}

func (m ExchangeLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
