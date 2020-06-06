// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsRemove struct{}

func (m ItemsRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsRemove
}

func (m ItemsRemove) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsRemove) Deserialize(extra string) error {
	return nil
}
