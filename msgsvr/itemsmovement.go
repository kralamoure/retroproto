// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsMovement struct{}

func (m ItemsMovement) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsMovement
}

func (m ItemsMovement) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsMovement) Deserialize(extra string) error {
	return nil
}
