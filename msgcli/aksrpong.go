// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AksRPong struct{}

func (m AksRPong) ProtocolId() d1proto.MsgCliId {
	return d1proto.AksRPong
}

func (m AksRPong) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AksRPong) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
