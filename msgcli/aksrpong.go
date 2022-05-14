// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AksRPong struct{}

func (m AksRPong) MessageId() retroproto.MsgCliId {
	return retroproto.AksRPong
}

func (m AksRPong) MessageName() string {
	return "AksRPong"
}

func (m AksRPong) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksRPong) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
