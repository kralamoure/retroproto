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
	return "", nil
}

func (m *KeyCreate) Deserialize(extra string) error {
	return nil
}
