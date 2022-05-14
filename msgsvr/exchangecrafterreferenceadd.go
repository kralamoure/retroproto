// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCrafterReferenceAdd struct{}

func (m ExchangeCrafterReferenceAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCrafterReferenceAdd
}

func (m ExchangeCrafterReferenceAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
