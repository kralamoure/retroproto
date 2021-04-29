// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeSetPublicMode struct{}

func (m ExchangeSetPublicMode) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeSetPublicMode
}

func (m ExchangeSetPublicMode) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeSetPublicMode) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
