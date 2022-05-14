// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCraftSuccess struct{}

func (m ExchangeCraftSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftSuccess
}

func (m ExchangeCraftSuccess) MessageName() string {
	return "ExchangeCraftSuccess"
}

func (m ExchangeCraftSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
