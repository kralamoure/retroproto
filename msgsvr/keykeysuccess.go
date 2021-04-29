// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type KeyKeySuccess struct{}

func (m KeyKeySuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.KeyKeySuccess
}

func (m KeyKeySuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *KeyKeySuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
