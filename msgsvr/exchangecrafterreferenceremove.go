// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCrafterReferenceRemove struct{}

func (m ExchangeCrafterReferenceRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCrafterReferenceRemove
}

func (m ExchangeCrafterReferenceRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
