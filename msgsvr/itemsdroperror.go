// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsDropError struct{}

func (m ItemsDropError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsDropError
}

func (m ItemsDropError) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsDropError) Deserialize(extra string) error {
	return nil
}
