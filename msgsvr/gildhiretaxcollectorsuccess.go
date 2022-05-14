// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildHireTaxCollectorSuccess struct{}

func (m GildHireTaxCollectorSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GildHireTaxCollectorSuccess
}

func (m GildHireTaxCollectorSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildHireTaxCollectorSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
