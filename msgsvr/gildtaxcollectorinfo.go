// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildTaxCollectorInfo struct{}

func (m GildTaxCollectorInfo) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildTaxCollectorInfo
}

func (m GildTaxCollectorInfo) Serialized() (string, error) {
	return "", nil
}

func (m *GildTaxCollectorInfo) Deserialize(extra string) error {
	return nil
}
