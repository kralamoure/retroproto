// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeSetPublicMode struct{}

func (m ExchangeSetPublicMode) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeSetPublicMode
}

func (m ExchangeSetPublicMode) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeSetPublicMode) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}