// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type InfosQuantity struct{}

func (m InfosQuantity) ProtocolId() d1proto.MsgSvrId {
	return d1proto.InfosQuantity
}

func (m InfosQuantity) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *InfosQuantity) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
