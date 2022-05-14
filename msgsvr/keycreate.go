// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type KeyCreate struct{}

func (m KeyCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyCreate
}

func (m KeyCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
