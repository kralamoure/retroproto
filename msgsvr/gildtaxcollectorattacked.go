// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildTaxCollectorAttacked struct{}

func (m GildTaxCollectorAttacked) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildTaxCollectorAttacked
}

func (m GildTaxCollectorAttacked) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GildTaxCollectorAttacked) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
