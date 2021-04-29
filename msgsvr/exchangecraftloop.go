// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCraftLoop struct{}

func (m ExchangeCraftLoop) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCraftLoop
}

func (m ExchangeCraftLoop) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCraftLoop) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
