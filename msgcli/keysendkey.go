// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type KeySendKey struct{}

func (m KeySendKey) ProtocolId() retroproto.MsgCliId {
	return retroproto.KeySendKey
}

func (m KeySendKey) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeySendKey) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
