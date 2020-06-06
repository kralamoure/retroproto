// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildHireTaxCollectorSuccess struct{}

func (m GildHireTaxCollectorSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildHireTaxCollectorSuccess
}

func (m GildHireTaxCollectorSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *GildHireTaxCollectorSuccess) Deserialize(extra string) error {
	return nil
}
