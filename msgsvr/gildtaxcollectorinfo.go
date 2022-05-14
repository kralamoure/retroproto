// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildTaxCollectorInfo struct{}

func (m GildTaxCollectorInfo) MessageId() retroproto.MsgSvrId {
	return retroproto.GildTaxCollectorInfo
}

func (m GildTaxCollectorInfo) MessageName() string {
	return "GildTaxCollectorInfo"
}

func (m GildTaxCollectorInfo) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildTaxCollectorInfo) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
