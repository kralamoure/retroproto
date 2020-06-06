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
	return "", nil
}

func (m *StoragesListRemove) Deserialize(extra string) error {
	return nil
}
