// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type KeyRequestLeave struct{}

func (m KeyRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.KeyRequestLeave
}

func (m KeyRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeyRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
