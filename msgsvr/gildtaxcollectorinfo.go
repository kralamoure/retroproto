// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildTaxCollectorInfo struct{}

func (m GildTaxCollectorInfo) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildTaxCollectorInfo
}

func (m GildTaxCollectorInfo) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildTaxCollectorInfo) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
