// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsItemUseCondition struct{}

func (m ItemsItemUseCondition) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsItemUseCondition
}

func (m ItemsItemUseCondition) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsItemUseCondition) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
