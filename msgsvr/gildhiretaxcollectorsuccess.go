// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildHireTaxCollectorSuccess struct{}

func (m GildHireTaxCollectorSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildHireTaxCollectorSuccess
}

func (m GildHireTaxCollectorSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildHireTaxCollectorSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
