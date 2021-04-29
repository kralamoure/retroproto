// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type StoragesLockedProperty struct{}

func (m StoragesLockedProperty) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.StoragesLockedProperty
}

func (m StoragesLockedProperty) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *StoragesLockedProperty) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
