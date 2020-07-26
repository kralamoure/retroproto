// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type StoragesListRemove struct{}

func (m StoragesListRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.StoragesListRemove
}

func (m StoragesListRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *StoragesListRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
