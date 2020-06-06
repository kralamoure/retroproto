// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsAddError struct{}

func (m ItemsAddError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsAddError
}

func (m ItemsAddError) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsAddError) Deserialize(extra string) error {
	return nil
}
