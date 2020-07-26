// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLocalDistantSuccess struct{}

func (m ExchangeLocalDistantSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLocalDistantSuccess
}

func (m ExchangeLocalDistantSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeLocalDistantSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
