// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsQuantity struct{}

func (m ItemsQuantity) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsQuantity
}

func (m ItemsQuantity) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsQuantity) Deserialize(extra string) error {
	return nil
}
