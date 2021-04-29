// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type StoragesListAdd struct{}

func (m StoragesListAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.StoragesListAdd
}

func (m StoragesListAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *StoragesListAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
