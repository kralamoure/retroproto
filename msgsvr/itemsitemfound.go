// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsItemFound struct{}

func (m ItemsItemFound) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsItemFound
}

func (m ItemsItemFound) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsItemFound) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
