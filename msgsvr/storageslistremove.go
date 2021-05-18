// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type StoragesListRemove struct{}

func (m StoragesListRemove) ProtocolId() retroproto.MsgSvrId {
	return retroproto.StoragesListRemove
}

func (m StoragesListRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesListRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
