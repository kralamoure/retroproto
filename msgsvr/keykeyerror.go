// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type KeyKeyError struct{}

func (m KeyKeyError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.KeyKeyError
}

func (m KeyKeyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyKeyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
