// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCraftError struct{}

func (m ExchangeCraftError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftError
}

func (m ExchangeCraftError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
