// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCrafterReferenceRemove struct{}

func (m ExchangeCrafterReferenceRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCrafterReferenceRemove
}

func (m ExchangeCrafterReferenceRemove) MessageName() string {
	return "ExchangeCrafterReferenceRemove"
}

func (m ExchangeCrafterReferenceRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
