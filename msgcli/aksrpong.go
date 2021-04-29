// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AksRPong struct{}

func (m AksRPong) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AksRPong
}

func (m AksRPong) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AksRPong) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
