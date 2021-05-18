// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeRequestSuccess struct{}

func (m ExchangeRequestSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeRequestSuccess
}

func (m ExchangeRequestSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
