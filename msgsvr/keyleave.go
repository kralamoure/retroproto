// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type KeyLeave struct{}

func (m KeyLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.KeyLeave
}

func (m KeyLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeyLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
