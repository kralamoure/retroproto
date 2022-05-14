// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type KeyLeave struct{}

func (m KeyLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyLeave
}

func (m KeyLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
