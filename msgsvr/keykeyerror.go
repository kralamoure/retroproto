// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type KeyKeyError struct{}

func (m KeyKeyError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.KeyKeyError
}

func (m KeyKeyError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeyKeyError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
