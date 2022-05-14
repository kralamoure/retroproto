// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type StoragesListAdd struct{}

func (m StoragesListAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.StoragesListAdd
}

func (m StoragesListAdd) MessageName() string {
	return "StoragesListAdd"
}

func (m StoragesListAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesListAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
