// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeList struct{}

func (m ExchangeList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeList
}

func (m ExchangeList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
