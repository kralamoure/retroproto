// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCraftLoopEnd struct{}

func (m ExchangeCraftLoopEnd) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftLoopEnd
}

func (m ExchangeCraftLoopEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftLoopEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
