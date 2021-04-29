// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCraftLoopEnd struct{}

func (m ExchangeCraftLoopEnd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCraftLoopEnd
}

func (m ExchangeCraftLoopEnd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCraftLoopEnd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
