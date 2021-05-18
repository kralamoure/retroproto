// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildTaxCollectorAttacked struct{}

func (m GildTaxCollectorAttacked) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GildTaxCollectorAttacked
}

func (m GildTaxCollectorAttacked) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildTaxCollectorAttacked) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
