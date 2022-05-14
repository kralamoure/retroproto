// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeList struct{}

func (m ExchangeList) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeList
}

func (m ExchangeList) MessageName() string {
	return "ExchangeList"
}

func (m ExchangeList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
