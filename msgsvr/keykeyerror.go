// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type KeyKeyError struct{}

func (m KeyKeyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.KeyKeyError
}

func (m KeyKeyError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *KeyKeyError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
