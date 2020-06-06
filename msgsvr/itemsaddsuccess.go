// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsAddSuccess struct{}

func (m ItemsAddSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsAddSuccess
}

func (m ItemsAddSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsAddSuccess) Deserialize(extra string) error {
	return nil
}
