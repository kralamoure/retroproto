// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type StoragesLockedProperty struct{}

func (m StoragesLockedProperty) ProtocolId() d1proto.MsgSvrId {
	return d1proto.StoragesLockedProperty
}

func (m StoragesLockedProperty) Serialized() (string, error) {
	return "", nil
}

func (m *StoragesLockedProperty) Deserialize(extra string) error {
	return nil
}
