// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCraftSuccess struct{}

func (m ExchangeCraftSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCraftSuccess
}

func (m ExchangeCraftSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCraftSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
