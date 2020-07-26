// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBuyError struct{}

func (m ExchangeBuyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBuyError
}

func (m ExchangeBuyError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeBuyError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
