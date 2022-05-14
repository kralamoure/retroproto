// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ItemsSetSkin struct{}

func (m ItemsSetSkin) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsSetSkin
}

func (m ItemsSetSkin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsSetSkin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
