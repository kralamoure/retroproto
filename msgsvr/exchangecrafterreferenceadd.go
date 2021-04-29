// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeCrafterReferenceAdd struct{}

func (m ExchangeCrafterReferenceAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCrafterReferenceAdd
}

func (m ExchangeCrafterReferenceAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
