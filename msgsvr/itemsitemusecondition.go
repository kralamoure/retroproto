// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsItemUseCondition struct{}

func (m ItemsItemUseCondition) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsItemUseCondition
}

func (m ItemsItemUseCondition) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsItemUseCondition) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
