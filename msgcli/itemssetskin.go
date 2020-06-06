// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsSetSkin struct{}

func (m ItemsSetSkin) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsSetSkin
}

func (m ItemsSetSkin) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsSetSkin) Deserialize(extra string) error {
	return nil
}
