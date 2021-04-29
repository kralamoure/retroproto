// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsSetSkin struct{}

func (m ItemsSetSkin) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ItemsSetSkin
}

func (m ItemsSetSkin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsSetSkin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
