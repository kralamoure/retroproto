// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ItemsItemFound struct{}

func (m ItemsItemFound) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsItemFound
}

func (m ItemsItemFound) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ItemsItemFound) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
