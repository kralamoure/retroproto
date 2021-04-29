// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type InfosQuantity struct{}

func (m InfosQuantity) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.InfosQuantity
}

func (m InfosQuantity) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *InfosQuantity) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
