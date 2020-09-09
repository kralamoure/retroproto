package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBuyError struct{}

func (m ExchangeBuyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBuyError
}

func (m ExchangeBuyError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuyError) Deserialize(extra string) error {
	return nil
}
