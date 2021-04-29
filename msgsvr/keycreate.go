// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type KeyCreate struct{}

func (m KeyCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.KeyCreate
}

func (m KeyCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeyCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
