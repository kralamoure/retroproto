// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsUseConfirm struct{}

func (m ItemsUseConfirm) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsUseConfirm
}

func (m ItemsUseConfirm) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsUseConfirm) Deserialize(extra string) error {
	return nil
}
