package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBuySuccess struct{}

func (m ExchangeBuySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBuySuccess
}

func (m ExchangeBuySuccess) MessageName() string {
	return "ExchangeBuySuccess"
}

func (m ExchangeBuySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuySuccess) Deserialize(extra string) error {
	return nil
}
