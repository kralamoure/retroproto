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
	return "", nil
}

func (m *GildTaxCollectorAttacked) Deserialize(extra string) error {
	return nil
}
