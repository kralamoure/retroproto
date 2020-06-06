// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsDropSuccess struct{}

func (m ItemsDropSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsDropSuccess
}

func (m ItemsDropSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsDropSuccess) Deserialize(extra string) error {
	return nil
}
