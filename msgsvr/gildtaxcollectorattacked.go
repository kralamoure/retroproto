// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildTaxCollectorAttacked struct{}

func (m GildTaxCollectorAttacked) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildTaxCollectorAttacked
}

func (m GildTaxCollectorAttacked) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildTaxCollectorAttacked) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
