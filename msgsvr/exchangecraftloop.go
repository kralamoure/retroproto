// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCraftLoop struct{}

func (m ExchangeCraftLoop) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCraftLoop
}

func (m ExchangeCraftLoop) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCraftLoop) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
