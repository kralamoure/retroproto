// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsUseNoConfirm struct{}

func (m ItemsUseNoConfirm) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsUseNoConfirm
}

func (m ItemsUseNoConfirm) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsUseNoConfirm) Deserialize(extra string) error {
	return nil
}
