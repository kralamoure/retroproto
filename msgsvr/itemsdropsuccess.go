// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsDropSuccess struct{}

func (m ItemsDropSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsDropSuccess
}

func (m ItemsDropSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDropSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
