// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type InfosQuantity struct{}

func (m InfosQuantity) ProtocolId() retroproto.MsgSvrId {
	return retroproto.InfosQuantity
}

func (m InfosQuantity) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosQuantity) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
