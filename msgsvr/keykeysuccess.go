// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type KeyKeySuccess struct{}

func (m KeyKeySuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.KeyKeySuccess
}

func (m KeyKeySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyKeySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
