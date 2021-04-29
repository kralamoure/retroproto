// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsDropSuccess struct{}

func (m ItemsDropSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsDropSuccess
}

func (m ItemsDropSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsDropSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
