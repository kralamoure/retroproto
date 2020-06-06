// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsAccessories struct{}

func (m ItemsAccessories) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsAccessories
}

func (m ItemsAccessories) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsAccessories) Deserialize(extra string) error {
	return nil
}
