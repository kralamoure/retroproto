package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBuyError struct{}

func (m ExchangeBuyError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBuyError
}

func (m ExchangeBuyError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuyError) Deserialize(extra string) error {
	return nil
}
