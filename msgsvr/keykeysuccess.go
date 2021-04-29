// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type KeyKeySuccess struct{}

func (m KeyKeySuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.KeyKeySuccess
}

func (m KeyKeySuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeyKeySuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
