// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildHireTaxCollectorError struct{}

func (m GildHireTaxCollectorError) MessageId() retroproto.MsgSvrId {
	return retroproto.GildHireTaxCollectorError
}

func (m GildHireTaxCollectorError) MessageName() string {
	return "GildHireTaxCollectorError"
}

func (m GildHireTaxCollectorError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildHireTaxCollectorError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
