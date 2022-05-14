// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeCraftLoop struct{}

func (m ExchangeCraftLoop) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftLoop
}

func (m ExchangeCraftLoop) MessageName() string {
	return "ExchangeCraftLoop"
}

func (m ExchangeCraftLoop) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftLoop) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
