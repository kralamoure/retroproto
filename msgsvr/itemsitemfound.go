// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsItemFound struct{}

func (m ItemsItemFound) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemFound
}

func (m ItemsItemFound) MessageName() string {
	return "ItemsItemFound"
}

func (m ItemsItemFound) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsItemFound) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
