// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type StoragesListRemove struct{}

func (m StoragesListRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.StoragesListRemove
}

func (m StoragesListRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *StoragesListRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
