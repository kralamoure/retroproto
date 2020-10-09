// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type KeyRequestLeave struct{}

func (m KeyRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.KeyRequestLeave
}

func (m KeyRequestLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *KeyRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
