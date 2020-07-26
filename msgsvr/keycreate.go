// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type KeyCreate struct{}

func (m KeyCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.KeyCreate
}

func (m KeyCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *KeyCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
