// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsItemUseCondition struct{}

func (m ItemsItemUseCondition) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsItemUseCondition
}

func (m ItemsItemUseCondition) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ItemsItemUseCondition) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
