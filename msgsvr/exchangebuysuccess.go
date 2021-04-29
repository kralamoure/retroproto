package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBuySuccess struct{}

func (m ExchangeBuySuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBuySuccess
}

func (m ExchangeBuySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuySuccess) Deserialize(extra string) error {
	return nil
}
