// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildHireTaxCollectorError struct{}

func (m GildHireTaxCollectorError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildHireTaxCollectorError
}

func (m GildHireTaxCollectorError) Serialized() (string, error) {
	return "", nil
}

func (m *GildHireTaxCollectorError) Deserialize(extra string) error {
	return nil
}
