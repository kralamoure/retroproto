// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type KeyRequestLeave struct{}

func (m KeyRequestLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.KeyRequestLeave
}

func (m KeyRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
