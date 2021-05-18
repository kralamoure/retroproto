// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type StoragesLockedProperty struct{}

func (m StoragesLockedProperty) ProtocolId() retroproto.MsgSvrId {
	return retroproto.StoragesLockedProperty
}

func (m StoragesLockedProperty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesLockedProperty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
