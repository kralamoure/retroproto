// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsRequestMovement struct{}

func (m ItemsRequestMovement) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsRequestMovement
}

func (m ItemsRequestMovement) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsRequestMovement) Deserialize(extra string) error {
	return nil
}
