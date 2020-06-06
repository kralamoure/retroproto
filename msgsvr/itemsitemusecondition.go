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
	return "", nil
}

func (m *ItemsItemUseCondition) Deserialize(extra string) error {
	return nil
}
