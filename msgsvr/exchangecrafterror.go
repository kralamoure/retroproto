// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCraftError struct{}

func (m ExchangeCraftError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCraftError
}

func (m ExchangeCraftError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCraftError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
