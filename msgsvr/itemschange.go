// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsChange struct{}

func (m ItemsChange) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsChange
}

func (m ItemsChange) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsChange) Deserialize(extra string) error {
	return nil
}
