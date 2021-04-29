// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCrafterReferenceRemove struct{}

func (m ExchangeCrafterReferenceRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCrafterReferenceRemove
}

func (m ExchangeCrafterReferenceRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
