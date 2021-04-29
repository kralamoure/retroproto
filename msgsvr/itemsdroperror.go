// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsDropError struct{}

func (m ItemsDropError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsDropError
}

func (m ItemsDropError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsDropError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
