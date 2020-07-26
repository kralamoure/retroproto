// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCraftSuccess struct{}

func (m ExchangeCraftSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCraftSuccess
}

func (m ExchangeCraftSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCraftSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
