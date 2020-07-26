// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCrafterReferenceAdd struct{}

func (m ExchangeCrafterReferenceAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCrafterReferenceAdd
}

func (m ExchangeCrafterReferenceAdd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
