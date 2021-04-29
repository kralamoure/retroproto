// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type StoragesListAdd struct{}

func (m StoragesListAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.StoragesListAdd
}

func (m StoragesListAdd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *StoragesListAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
