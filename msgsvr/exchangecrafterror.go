// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCraftError struct{}

func (m ExchangeCraftError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCraftError
}

func (m ExchangeCraftError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCraftError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
