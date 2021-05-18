// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsChange struct{}

func (m ItemsChange) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsChange
}

func (m ItemsChange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsChange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
