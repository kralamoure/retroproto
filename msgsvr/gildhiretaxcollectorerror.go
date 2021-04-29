// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildHireTaxCollectorError struct{}

func (m GildHireTaxCollectorError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildHireTaxCollectorError
}

func (m GildHireTaxCollectorError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildHireTaxCollectorError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
