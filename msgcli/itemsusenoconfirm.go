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
	return "", d1proto.ErrNotImplemented
}

func (m *ItemsUseNoConfirm) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
