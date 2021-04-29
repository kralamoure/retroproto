// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsChange struct{}

func (m ItemsChange) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsChange
}

func (m ItemsChange) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsChange) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
